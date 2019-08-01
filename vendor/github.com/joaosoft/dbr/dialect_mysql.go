package dbr

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"time"
)

type dialectMySql struct{}

func (d *dialectMySql) Name() string {
	return string(constDialectMysql)
}

func (d *dialectMySql) Encode(i interface{}) string {
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

// https://dev.mysql.com/doc/refman/5.7/en/string-literals.html
func (d *dialectMySql) EncodeString(s string) string {
	buf := new(bytes.Buffer)

	buf.WriteRune('\'')
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 0:
			buf.WriteString(`\0`)
		case '\'':
			buf.WriteString(`\'`)
		case '"':
			buf.WriteString(`\"`)
		case '\b':
			buf.WriteString(`\b`)
		case '\n':
			buf.WriteString(`\n`)
		case '\r':
			buf.WriteString(`\r`)
		case '\t':
			buf.WriteString(`\t`)
		case 26:
			buf.WriteString(`\Z`)
		case '\\':
			buf.WriteString(`\\`)
		default:
			buf.WriteByte(s[i])
		}
	}

	buf.WriteRune('\'')
	return buf.String()
}

func (d *dialectMySql) EncodeBool(b bool) string {
	if b {
		return constMySqlBoolTrue
	}
	return constMySqlBoolFalse
}

func (d *dialectMySql) EncodeTime(t time.Time) string {
	return `'` + t.UTC().Format(constTimeFormat) + `'`
}

func (d *dialectMySql) EncodeBytes(b []byte) string {
	return fmt.Sprintf(`0x%x`, b)
}

func (d *dialectMySql) EncodeColumn(column interface{}) string {
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

func (d *dialectMySql) Placeholder() string {
	return constMysqlPlaceHolder
}
