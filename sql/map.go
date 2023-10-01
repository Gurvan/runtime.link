package sql

import (
	"context"
	"reflect"
	"strconv"
	"strings"
	"sync"

	"runtime.link/sql/std/sodium"
)

// Map represents a distinct mapping of data stored in a [Database].
type Map[K comparable, V any] struct {
	to sodium.Table
	db sodium.Database
}

// Open a new [Map] from the given [Database]. The table schema is
// derived from the key and value types 'K' and 'V', following the
// same rules as [ValuesOf]. A 'sql' or else, a 'txt' tag controls
// the name of the column. If no tag is specified, the ToLower(name)
// of the field is used. If the key is not a struct, the column name
// is 'id', otherwise it is treated as a composite key across each
// field. If the value is not a struct, the column name is 'value'.
// Nested structures are named with an underscore used to seperate
// the field path unless the structure is embedded, in which case
// the nested fields are promoted. Arrays elements are suffixed by
// their index.
func Open[K comparable, V any](db sodium.Database, table string) Map[K, V] {
	key := reflect.StructField{
		Name: "id",
		Type: reflect.TypeOf([0]K{}).Elem(),
	}
	if key.Type.Kind() == reflect.Struct {
		key.Anonymous = true
	}
	val := reflect.StructField{
		Name: "value",
		Type: reflect.TypeOf([0]V{}).Elem(),
	}
	if val.Type.Kind() == reflect.Struct {
		val.Anonymous = true
	}
	sentinals.index.assert(key, new(K))
	sentinals.value.assert(val, new(V))
	return Map[K, V]{
		to: sodium.Table{
			Name:  table,
			Index: columnsOf(key),
			Value: columnsOf(val),
		},
		db: db,
	}
}

// OpenTable a new [Map] from the given [Database] and specified
// table.
func OpenTable[K comparable, V any](db sodium.Database, table sodium.Table) Map[K, V] {
	return Map[K, V]{
		to: table,
		db: db,
	}
}

// Insert a new value into the [Map] at the given key. The given [Flag] determines
// how the value is inserted. If the [Flag] is [Upsert], the value will overwrite
// any existing value at the given key. If the [Flag] is [Create], the value will
// only be inserted if there is no existing value at the given key, otherwise an
// error will be returned.
func (m Map[K, V]) Insert(ctx context.Context, key K, flag Flag, value V) error {
	tx, err := m.db.Manage(ctx, 0)
	if err != nil {
		return err
	}
	insert := m.db.Insert(m.to, ValuesOf(key), bool(flag), ValuesOf(value))
	select {
	case tx <- insert:
	case <-ctx.Done():
		return ctx.Err()
	}
	close(tx)
	n, err := insert.Wait(ctx)
	if err != nil {
		return err
	}
	if n == -1 {
		return ErrDuplicate
	}
	return nil
}

func (m Map[K, V]) SearchFunc(ctx context.Context, query QueryFunc[K, V]) Chan[K, V] {
	key := sentinals.index[reflect.TypeOf([0]K{}).Elem()].(*K)
	val := sentinals.value[reflect.TypeOf([0]V{}).Elem()].(*V)
	sql := query(key, val)
	out := make(Chan[K, V])
	go func() {
		defer close(out)
		ch := make(chan []sodium.Value, 64)
		do := m.db.Search(m.to, sodium.Query(sql), ch)
		tx, err := m.db.Manage(ctx, 0)
		if err != nil {
			select {
			case out <- result[K, V]{err: err}:
				return
			case <-ctx.Done():
				return
			}
		}
		select {
		case tx <- do:
		case <-ctx.Done():
			return
		}
		for values := range ch {
			var result result[K, V]
			decode(reflect.ValueOf(&result.key), values[:len(m.to.Index)])
			decode(reflect.ValueOf(&result.val), values[len(m.to.Index):])
			select {
			case out <- result:
				continue
			case <-ctx.Done():
				return
			}
		}
		close(tx)
	}()
	return out
}

var smutex sync.RWMutex
var mirror = make(map[any][]sodium.Column)

var sentinals struct {
	index sentinal
	value sentinal
}

type sentinal map[reflect.Type]any

func (s *sentinal) assert(field reflect.StructField, arg any) {
	smutex.Lock()
	defer smutex.Unlock()
	if *s == nil {
		*s = make(map[reflect.Type]any)
	}
	_, ok := (*s)[field.Type]
	if ok {
		return
	}
	(*s)[field.Type] = arg
	s.walk(field, reflect.ValueOf(arg).Elem())
}

func (s *sentinal) walk(field reflect.StructField, arg reflect.Value, path ...string) {
	name := strings.ToLower(field.Name)
	if tag := field.Tag.Get("txt"); tag != "" {
		name = tag
	}
	if tag := field.Tag.Get("sql"); tag != "" {
		name = tag
	}
	if len(path) > 0 {
		name = strings.Join(path, "_") + "_" + name
	}
	mirror[arg.Addr().Interface()] = columnsOf(field, path...)
	switch field.Type.Kind() {
	case reflect.Struct:
		for i := 0; i < field.Type.NumField(); i++ {
			promote := append(path, name)
			if field.Type.Field(i).Anonymous {
				promote = path
			}
			s.walk(field.Type.Field(i), arg.Field(i), promote...)
		}
	case reflect.Array:
		for i := 0; i < field.Type.Len(); i++ {
			s.walk(field.Type.Field(i), arg.Index(i), append(path, name+strconv.Itoa(i))...)
		}
	}
}
