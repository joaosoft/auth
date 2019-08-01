package dbr

import (
	"fmt"
)

type caseWhens []*caseWhen

func newCaseWhens() caseWhens {
	return make(caseWhens, 0)
}

func (c caseWhens) Build(db *db) (string, error) {
	var query string

	if len(c) == 0 {
		return "", nil
	}

	for i, cond := range c {
		onWhen, err := cond.Build(db)
		if err != nil {
			return "", err
		}

		if i > 0 {
			query += " "
		}

		query += fmt.Sprintf("%s", onWhen)
	}

	return query, nil
}
