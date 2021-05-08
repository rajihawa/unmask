package utils

import (
	"database/sql"
	"errors"
	"reflect"
)

func ScanStruct(r *sql.Rows, dest interface{}) error {
	v := reflect.ValueOf(dest)

	if v.Kind() != reflect.Ptr {
		return errors.New("must pass a pointer, not a value, to ScanStruct dest")
	}

	v = v.Elem()

	columns, err := r.Columns()
	if err != nil {
		return err
	}
	colLen := len(columns)
	values := make([]interface{}, colLen)

	v = reflect.Indirect(v)
	if v.Kind() != reflect.Struct {
		return errors.New("argument not a struct")
	}
	for i := 0; i < colLen; i++ {
		f := v.FieldByIndex([]int{i})
		values[i] = f.Addr().Interface()
	}
	return r.Scan(values...)
}
