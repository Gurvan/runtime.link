package txt

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type (
	MayPrefix  bool
	MayContain string
)

func (p MayPrefix) String() string { return strconv.FormatBool(bool(p)) }

func (p MayPrefix) MatchString(ptr any, raw string, tag reflect.StructTag) (n int, err error) {
	if strings.HasPrefix(raw, string(tag)) {
		*(ptr.(*Prefix)) = true
		return len(tag), nil
	}
	*(ptr.(*Prefix)) = false
	return 0, nil
}

func (c MayContain) String() string { return string(c) }
func (MayContain) MatchString(ptr any, raw string, tag reflect.StructTag) (n int, err error) {
	contains := ptr.(*MayContain)
	for _, char := range raw {
		if !strings.ContainsRune(string(tag), char) {
			return 0, fmt.Errorf("invalid '%v' character", string(char))
		}
		*contains += MayContain(char)
	}
	return 0, nil
}

type WithBacktick[T Matcher] struct {
	WithBacktick T
}

func (b WithBacktick[T]) String() string { return b.WithBacktick.String() }
func (WithBacktick[T]) MatchString(ptr any, raw string, tag reflect.StructTag) (n int, err error) {
	val := *(ptr.(*WithBacktick[T]))
	return val.WithBacktick.MatchString(&val.WithBacktick, raw, tag+"`")
}
