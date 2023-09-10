package cmd

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"reflect"
	"strconv"
	"strings"
	"syscall"

	"runtime.link/std"
)

// Import the given program, using the specified name (which should be a file or in the system PATH).
func Import[Program any](names ...string) Program {
	var program Program
	structure := std.StructureOf(&program)
	var found bool
	for _, name := range names {
		_, err := exec.LookPath(name)
		if err == nil {
			set(structure, name)
			found = true
			break
		}
	}
	if host := structure.Host.Get("cmd"); host != "" {
		_, err := exec.LookPath(host)
		if err == nil {
			set(structure, host)
			found = true
			return program
		}
	}
	if !found {
		structure.MakeError(errors.New("cannot find program: " + strings.Join(names, ", ")))
	}
	return program
}

type listArguments []string

func (execArgs *listArguments) add(val reflect.Value) error {
	switch val.Kind() {
	case reflect.Struct:
		rtype := val.Type()
		for i := 0; i < rtype.NumField(); i++ {
			field := rtype.Field(i)
			if !field.IsExported() {
				continue
			}
			if field.Anonymous && field.Type.Kind() == reflect.Struct {
				if err := execArgs.add(val.Field(i)); err != nil {
					return err
				}
				continue
			}
			exec := field.Tag.Get("cmd")
			if exec == "-" {
				continue
			}

			omitBooleanFlag := (!strings.Contains(exec, "%v") && field.Type.Kind() == reflect.Bool)

			if (strings.Contains(exec, ",omitempty") || omitBooleanFlag) && val.Field(i).IsZero() {
				continue
			}

			exec, _, _ = strings.Cut(exec, ",")

			parts := strings.Split(exec, " ")
			for _, part := range parts {
				if strings.HasPrefix(part, "%") || strings.Contains(part, "%") {
					*execArgs = append(*execArgs, fmt.Sprintf(part, val.Field(i).Interface()))
				} else {
					*execArgs = append(*execArgs, part)
				}
			}
		}
	case reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			if err := execArgs.add(val.Index(i)); err != nil {
				return err
			}
		}
	case reflect.Pointer:
		if !val.IsNil() {
			if err := execArgs.add(val.Elem()); err != nil {
				return err
			}
		}
	default:
		*execArgs = append(*execArgs, fmt.Sprint(val.Interface()))
	}
	return nil
}

func set(spec std.Structure, cmd string) {
	for _, fn := range spec.Functions {
		fn := fn
		tag := string(fn.Tags.Get("cmd"))

		cmd := cmd
		if cmd == "" {
			cmd, tag, _ = strings.Cut(tag, " ")
		}

		var isJSON bool = false
		if strings.HasSuffix(tag, " | json") {
			tag = strings.TrimSuffix(tag, " | json")
			isJSON = true
		}

		fn.Make(reflect.MakeFunc(fn.Type, func(args []reflect.Value) (results []reflect.Value) {
			ctx := context.Background()
			if fn.Type.In(0) == reflect.TypeOf([0]context.Context{}).Elem() {
				ctx = args[0].Interface().(context.Context)
				args = args[1:]
			}

			scanner := std.NewArgumentScanner(args)

			var execArgs listArguments
			if tag != "" {
				for _, component := range strings.Split(string(tag), " ") {
					if strings.HasPrefix(component, "%") || strings.HasPrefix(component, "{") {
						component = strings.Trim(component, "{}")
						val, err := scanner.Scan(component)
						if err != nil {
							return fn.Return(nil, err)
						}
						if err := execArgs.add(val); err != nil {
							return fn.Return(nil, err)
						}
					} else {
						execArgs = append(execArgs, component)
					}
				}
			}

			var stdout bytes.Buffer
			var stderr bytes.Buffer

			stdoutRead, stdoutWrite, err := os.Pipe()
			if err != nil {
				return fn.Return(nil, err)
			}
			stderrRead, stderrWrite, err := os.Pipe()
			if err != nil {
				return fn.Return(nil, err)
			}
			if os.Getenv("DEBUG_CMD") != "" {
				fmt.Println(cmd, execArgs)
			}
			cmd := exec.CommandContext(ctx, cmd, execArgs...)
			cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
			cmd.Cancel = func() error {
				if err := stdoutWrite.Close(); err != nil {
					return err
				}
				if err := stderrWrite.Close(); err != nil {
					return err
				}
				if err := syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL); err != nil {
					return cmd.Process.Kill()
				}
				return nil
			}

			results = make([]reflect.Value, fn.Type.NumOut())
			for i := range results {
				results[i] = reflect.Zero(fn.Type.Out(i))
			}

			var async bool
			var chout chan []byte
			var cherr chan error

			if fn.NumOut() > 0 && fn.Type.Out(0) == reflect.TypeOf([0]chan []byte{}).Elem() {
				cmd.Stdout = stdoutWrite
				async = true
				chout = make(chan []byte)
				results[0] = reflect.ValueOf(chout)
				go func() {
					reader := bufio.NewReader(stdoutRead)
					for {
						line, err := reader.ReadBytes('\n')
						if err != nil {
							if err != io.EOF {

							}
							close(chout)
							return
						}
						select {
						case chout <- line[:len(line)-1]:
						case <-ctx.Done():
							return
						}
					}
				}()
			} else if fn.NumOut() > 0 {
				cmd.Stdout = stdoutWrite
				go io.Copy(&stdout, stdoutRead)
			} else {
				cmd.Stdout = os.Stdout
			}

			if fn.Type.NumOut() > 0 && fn.Type.Out(fn.Type.NumOut()-1) == reflect.TypeOf([0]chan error{}).Elem() {
				cmd.Stderr = stderrWrite
				async = true
				cherr = make(chan error)
				results[1] = reflect.ValueOf(cherr)
				go func() {
					reader := bufio.NewReader(stderrRead)
					for {
						line, err := reader.ReadBytes('\n')
						if err != nil {
							if err != io.EOF {

							}
							close(cherr)
							return
						}
						select {
						case cherr <- errors.New(string(line[:len(line)-1])):
						case <-ctx.Done():
							return
						}
					}
				}()
			} else if fn.NumOut() != fn.Type.NumOut() {
				cmd.Stderr = stderrWrite
				go io.Copy(&stderr, stderrRead)
			} else {
				cmd.Stderr = os.Stderr
			}
			if async {
				if err := cmd.Start(); err != nil {
					return fn.Return(results, err)
				}
				go func() {
					if err := cmd.Wait(); err != nil {
						select {
						case cherr <- err:
						case <-ctx.Done():
						}
					}
					stderrWrite.Close()
					stdoutWrite.Close()
				}()
				return
			}
			if err := cmd.Run(); err != nil {
				if text := stderr.String(); strings.TrimSpace(text) != "" {
					return fn.Return(nil, errors.New(text))
				}
				return fn.Return(nil, err)
			}
			if fn.NumOut() > 0 {
				if isJSON {
					var result = reflect.New(fn.Type.Out(0)).Interface()
					if err := json.NewDecoder(&stdout).Decode(result); err != nil {
						return fn.Return(nil, err)
					}
					return []reflect.Value{reflect.ValueOf(result).Elem()}
				} else {
					if fn.NumOut() == 1 {
						var value = reflect.New(fn.Type.Out(0)).Elem()
						switch fn.Type.Out(0).Kind() {
						case reflect.String:
							result := stdout.String()
							result = strings.TrimSuffix(result, "\n")
							results[0] = reflect.ValueOf(result)
							return fn.Return(results, nil)
						case reflect.Int32:
							result := stdout.String()
							result = strings.TrimSuffix(result, "\n")
							i, err := strconv.Atoi(result)
							if err != nil {
								return fn.Return(nil, err)
							}
							value.SetInt(int64(i))
							results[0] = value
							return fn.Return(results, nil)
						}
					}
					return fn.Return(nil, fmt.Errorf("exec: return type %v: not implemented", fn.Type.Out(0)))
				}
			}
			return
		}))
	}
	for _, section := range spec.Namespace {
		set(section, cmd)
	}
}