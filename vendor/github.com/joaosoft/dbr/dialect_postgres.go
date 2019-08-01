package dbr

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

type dialectPostgres struct{}

func (d *dialectPostgres) Name() string {
	return string(constDialectPostgres)
}

func (d *dialectPostgres) Encode(i interface{}) string {
	value := reflect.ValueOf(i)

	if value.Kind() == reflect.Ptr {
		if value.IsNil() {
			return constFunctionNull
		}
		value = value.Elem()
	}

	switch value.Kind() {
	case reflect.String:
		return d.EncodeString(value.String())
	case reflect.Bool:
		return d.EncodeBool(value.Bool())
	default:
		switch value.Type() {
		case reflect.TypeOf(time.Time{}):
			return d.EncodeTime(i.(time.Time))
		case reflect.TypeOf([]byte{}):
			return d.EncodeBytes(i.([]byte))
		}
	}

	return fmt.Sprintf("%+v", value.Interface())
}

func (d *dialectPostgres) EncodeString(s string) string {
	return `'` + strings.Replace(s, `'`, `''`, -1) + `'`
}

func (d *dialectPostgres) EncodeBool(b bool) string {
	if b {
		return constPostgresBoolTrue
	}
	return constPostgresBoolFalse
}

func (d *dialectPostgres) EncodeTime(t time.Time) string {
	return `'` + t.UTC().Format(constTimeFormat) + `'`
}

func (d *dialectPostgres) EncodeBytes(b []byte) string {
	return fmt.Sprintf(`E'\\x%x'`, b)
}

func (d *dialectPostgres) EncodeColumn(column interface{}) string {
	value := fmt.Sprintf("%+v", column)

	switch column.(type) {
	case string:
		if !strings.ContainsAny(value, `*`) {
			value = fmt.Sprintf(`"%s"`, value)
			value = strings.Replace(value, `.`, `"."`, 1)
		}
	}

	return value
}

func (d *dialectPostgres) Placeholder() string {
	return constPostgresPlaceHolder
}
