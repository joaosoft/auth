package dbr

import (
	"fmt"
	"strings"
)

type functionCase struct {
	alias   *string
	onWhens caseWhens
	onElse  *caseElse

	*functionBase
}

func newFunctionCase(alias ...string) *functionCase {
	funcCase := &functionCase{functionBase: newFunctionBase(false, false), onWhens: newCaseWhens()}

	if len(alias) > 0 {
		funcCase.alias = &alias[0]
	}

	return funcCase
}

func (c *functionCase) When(query string, values ...interface{}) *functionCase {
	c.onWhens = append(c.onWhens, newCaseWhen(newCondition(nil, OperatorAnd, query, values...)))

	return c
}

func (c *functionCase) Then(result interface{}) *functionCase {
	if len(c.onWhens) > 0 {
		c.onWhens[len(c.onWhens)-1].result = result
	}

	return c
}

func (c *functionCase) Else(result interface{}) *functionCase {
	c.onElse = newCaseElse(result)

	return c
}

func (c *functionCase) Expression(db *db) (string, error) {
	c.db = db
	return "", nil
}

func (c *functionCase) Build(db *db) (string, error) {
	c.db = db

	var value string
	var query string

	onWhens, err := c.onWhens.Build(db)
	if err != nil {
		return "", err
	}
	value += onWhens

	onElse, err := c.onElse.Build(db)
	if err != nil {
		return "", err
	}

	if len(onElse) > 0 {
		value += fmt.Sprintf(" %s", onElse)
	}

	query = fmt.Sprintf("(%s %s %s)", constFunctionCase, value, constFunctionEnd)

	if c.alias != nil && len(strings.TrimSpace(*c.alias)) > 0 {
		query += fmt.Sprintf(" %s %s", constFunctionAs, *c.alias)
	}

	return query, nil
}
