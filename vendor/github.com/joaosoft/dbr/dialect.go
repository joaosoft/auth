package dbr

import "time"

type DialectName string

var (
	DialectPostgres = &dialectPostgres{}
	DialectMySql    = &dialectMySql{}
	DialectSqlLite3 = &dialectSqlLite3{}
)

type dialect interface {
	Name() string
	Encode(i interface{}) string
	EncodeString(s string) string
	EncodeBool(b bool) string
	EncodeTime(t time.Time) string
	EncodeBytes(b []byte) string
	EncodeColumn(column interface{}) string
	Placeholder() string
}

func NewDialect(name string) dialect {
	switch name {
	case string(constDialectPostgres):
		return DialectPostgres
	case string(constDialectMysql):
		return DialectMySql
	case string(constDialectSqlLite3):
		return DialectSqlLite3
	}

	return nil
}
