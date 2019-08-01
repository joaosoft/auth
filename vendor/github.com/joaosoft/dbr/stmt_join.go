package dbr

import (
	"fmt"
)

type StmtJoin struct {
	join  Join
	table *table
	on    *condition

	db *db
}

type Join string

func newStmtJoin(db *db, join Join, table *table, on *condition) *StmtJoin {
	return &StmtJoin{
		db:    db,
		join:  join,
		table: table,
		on:    on,
	}
}

func (stmt *StmtJoin) Build() (string, error) {
	condition, err := stmt.on.Build()
	if err != nil {
		return "", err
	}

	table, err := stmt.table.Build()
	if err != nil {
		return "", err
	}

	query := fmt.Sprintf("%s %s ON (%s)", stmt.join, table, condition)

	return query, nil
}
