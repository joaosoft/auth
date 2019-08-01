package dbr

import (
	"database/sql"
	"fmt"
	"reflect"
	"time"
)

type StmtDelete struct {
	withStmt   *StmtWith
	table      *table
	conditions *conditions
	returning  *columns

	Dbr          *Dbr
	Db           *db
	Duration     time.Duration
	sqlOperation SqlOperation
}

func newStmtDelete(dbr *Dbr, db *db, withStmt *StmtWith) *StmtDelete {
	return &StmtDelete{
		sqlOperation: DeleteOperation,
		Dbr:          dbr,
		Db:           db,
		withStmt:     withStmt,
		conditions:   newConditions(dbr.Connections.Write),
		returning:    newColumns(db, false),
	}
}

func (stmt *StmtDelete) From(table interface{}) *StmtDelete {
	stmt.table = newTable(stmt.Db, table)
	return stmt
}

func (stmt *StmtDelete) Where(query interface{}, values ...interface{}) *StmtDelete {
	stmt.conditions.list = append(stmt.conditions.list, &condition{operator: OperatorAnd, query: query, values: values, db: stmt.Db})
	return stmt
}

func (stmt *StmtDelete) WhereOr(query string, values ...interface{}) *StmtDelete {
	stmt.conditions.list = append(stmt.conditions.list, &condition{operator: OperatorOr, query: query, values: values, db: stmt.Db})
	return stmt
}

func (stmt *StmtDelete) Build() (string, error) {
	var query string

	// withStmt
	if len(stmt.withStmt.withs) > 0 {
		withStmt, err := stmt.withStmt.Build()
		if err != nil {
			return "", err
		}
		query += withStmt
	}

	table, err := stmt.table.Build()
	if err != nil {
		return "", err
	}

	query += fmt.Sprintf("%s %s %s", constFunctionDelete, constFunctionFrom, table)

	if len(stmt.conditions.list) > 0 {
		conds, err := stmt.conditions.Build()
		if err != nil {
			return "", err
		}

		query += fmt.Sprintf(" %s %s", constFunctionWhere, conds)
	}

	if len(stmt.returning.list) > 0 {
		returning, err := stmt.returning.Build()
		if err != nil {
			return "", err
		}

		query += fmt.Sprintf(" %s %s", constFunctionReturning, returning)
	}

	return query, nil
}

func (stmt *StmtDelete) Exec() (sql.Result, error) {

	startTime := time.Now()
	defer func() {
		stmt.Duration = time.Since(startTime)
	}()

	query, err := stmt.Build()
	if err != nil {
		return nil, err
	}

	result, err := stmt.Db.Exec(query)

	table, err := stmt.table.Build()
	if err != nil {
		return nil, err
	}

	if err := stmt.Dbr.eventHandler(stmt.sqlOperation, []string{table}, query, err, nil, result); err != nil {
		return nil, err
	}

	return result, err
}

func (stmt *StmtDelete) Return(column ...interface{}) *StmtDelete {
	stmt.returning.list = append(stmt.returning.list, column...)
	return stmt
}

func (stmt *StmtDelete) Load(object interface{}) error {

	value := reflect.ValueOf(object)
	if value.Kind() != reflect.Ptr || value.IsNil() {
		return ErrorInvalidPointer
	}

	startTime := time.Now()
	defer func() {
		stmt.Duration = time.Since(startTime)
	}()

	query, err := stmt.Build()
	if err != nil {
		return err
	}

	rows, err := stmt.Db.Query(query)
	if err != nil {
		return err
	}

	table, err := stmt.table.Build()
	if err != nil {
		return err
	}
	if err := stmt.Dbr.eventHandler(stmt.sqlOperation, []string{table}, query, err, rows, nil); err != nil {
		return err
	}

	defer rows.Close()

	_, err = read(stmt.returning.list, rows, value)

	return err
}
