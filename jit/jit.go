/*
Package jit provides a safe alternative to [reflect.MakeFunc] with support for transparent optimisation.

This package is still in an experimental proof-of-concept phase and is not quite
ready for general use. The aim is to provide a safe way to create small optimised
functions at runtime. This package is included in runtime.link to serve as an
optimisation pathway for [api.Linker] implementations.
*/
package jit

import (
	"errors"
	"reflect"
	"unsafe"

	"runtime.link/bin"
	"runtime.link/bin/std/cpu"
)

// Implementation for a function.
type Implementation func(Assembly, []Value) ([]Value, error)

// Make a new function of type 'T' using the given JIT implementation
// function. The implementation function must have the JIT equivalent for
// each argument and return value. The Implementation function may be
// called each time the JIT function is called, in order to compile it,
// or it may only be called once. Therefore the behaviour of the
// implementation should not depend on any side effects or mutability.
func Make[T any](impl Implementation) (T, error) {
	var zero T
	val, err := MakeFunc(reflect.TypeOf([0]T{}).Elem(), impl)
	if err != nil {
		return zero, err
	}
	return val.Interface().(T), nil
}

// MakeFunc is like [Make], but it can be used to create a function value
// from a [reflect.Type] instead of one known at compile time.
func MakeFunc(fnType reflect.Type, impl Implementation) (reflect.Value, error) {
	var asm = Assembly{
		direct: true,
		fnType: fnType,
	}
	return reflect.MakeFunc(fnType, func(args []reflect.Value) (results []reflect.Value) {
		values := make([]Value, len(args))
		for i, arg := range args {
			values[i] = Value{direct: arg}
		}
		rets, err := impl(asm, values)
		if err != nil {
			panic(err)
		}
		results = make([]reflect.Value, len(rets))
		for i, ret := range rets {
			results[i] = ret.direct
		}
		return results
	}), nil
}

// Assembly used to assemble a function.
type Assembly struct {
	direct bool
	fnType reflect.Type // function being compiled.

	binary bin.Format
	gprs   []cpu.GPR
	fprs   []cpu.FPR
}

// ABI describes how to call an [unsafe.Pointer] as the specified function type.
type ABI interface {
	// Call the given pointer with the given arguments and return the results.
	// This will be used when making a 'safe' function using [reflect.MakeFunc] or
	// when the platform does not have JIT optimisations enabled. If an error is
	// returned, it will be returned by [MakeFunc]. The second return value should
	// contains functions that should release any resources allocated by the
	// corresponding argument.
	Call(unsafe.Pointer, []reflect.Value, ...reflect.Type) ([]reflect.Value, []func(), error)

	// CallingConvention returns the calling convention for the given function type
	// so that the location of each argument and return value is clearly specified,
	// if an error is returned, [Call] is used as a fallback.
	CallingConvention(reflect.Type) (args, rets []Location, err error)
}

// Add returns (a + b).
func (asm Assembly) Add(a, b Value) Value {
	if asm.direct {
		switch a.direct.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return Value{direct: reflect.ValueOf(a.direct.Int() + b.direct.Int()).Convert(a.direct.Type())}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			return Value{direct: reflect.ValueOf(a.direct.Uint() + b.direct.Uint()).Convert(a.direct.Type())}
		case reflect.Float32, reflect.Float64:
			return Value{direct: reflect.ValueOf(a.direct.Float() + b.direct.Float()).Convert(a.direct.Type())}
		default:
			panic("invalid value")
		}
	}
	return Value{}
}

// UnsafeCall uses the given ABI to call the given function pointer with the given arguments and
// return the results.
func (asm Assembly) UnsafeCall(abi ABI, ptr unsafe.Pointer, args []Value, rets ...reflect.Type) ([]Value, []Lifetime, error) {
	if !asm.direct {
		return nil, nil, errors.New("UnsafeCall is only available in direct mode")
	}
	var values = make([]reflect.Value, len(args))
	for i, arg := range args {
		values[i] = arg.direct
	}
	locals, freedom, err := abi.Call(ptr, values, rets...)
	if err != nil {
		for _, free := range freedom {
			if free != nil {
				free()
			}
		}
		return nil, nil, err
	}
	var lifetimes = make([]Lifetime, len(freedom))
	for i, free := range freedom {
		lifetimes[i] = Lifetime{direct: free}
	}
	var results = make([]Value, len(locals))
	for i, local := range locals {
		results[i] = Value{direct: local}
	}
	return results, lifetimes, nil
}
