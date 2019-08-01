package dbr

import (
	"fmt"
)

type caseElse struct {
	result interface{}
}

func newCaseElse(result interface{}) *caseElse {
	return &caseElse{result: result}
}

func (c *caseElse) Build(db *db) (string, error) {
	var query string
	var err error
	var result string

	// result
	switch stmt := c.result.(type) {
	case *StmtSelect:
		result, err = stmt.Build()
		if err != nil {
			return "", err
		}
		result = fmt.Sprintf("(%s)", result)
	default:
		if impl, ok := stmt.(ifunction); ok {
			result, err = impl.Build(db)
			if err != nil {
				return "", err
			}
		} else {
			result = fmt.Sprintf("%+v", stmt)
		}
	}

	query = fmt.Sprintf("%s %s", constFunctionElse, result)

	return query, nil
}
