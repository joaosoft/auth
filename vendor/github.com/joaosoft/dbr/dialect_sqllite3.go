package dbr

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

type dialectSqlLite3 struct{}

func (d *dialectSqlLite3) Name() string {
	return string(constDialectSqlLite3)
}

func (d *dialectSqlLite3) Encode(i interface{}) string {
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

// https://www.sqlite.org/faq.html
func (d *dialectSqlLite3) EncodeString(s string) string {
	return `'` + strings.Replace(s, `'`, `''`, -1) + `'`
}

// https://www.sqlite.org/lang_expr.html
func (d *dialectSqlLite3) EncodeBool(b bool) string {
	if b {
		return constSqlLite3BoolTrue
	}
	return constSqlLite3BoolFalse
}

// https://www.sqlite.org/lang_datefunc.html
func (d *dialectSqlLite3) EncodeTime(t time.Time) string {
	return `'` + t.UTC().Format(constTimeFormat) + `'`
}

// https://www.sqlite.org/lang_expr.html
func (d *dialectSqlLite3) EncodeBytes(b []byte) string {
	return fmt.Sprintf(`X'%x'`, b)
}

func (d *dialectSqlLite3) EncodeColumn(column interface{}) string {
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

func (d *dialectSqlLite3) Placeholder() string {
	return constSqlLite3PlaceHolder
}
