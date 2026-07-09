package storage

import (
	"context"
	"database/sql"
)

type Queries interface {
	// params: procName
	//
	// return: SQL query and args
	GetSQLQuery(string) (string, []any)
}

type Database struct {
	db *sql.DB
}

func NewDatabase(db *sql.DB) *Database {
	return &Database{db: db}
}

func (d *Database) GetDB() *sql.DB {
	return d.db
}

func (d *Database) Close() error {
	return d.db.Close()
}

func (d *Database) ExecQuery(ctx context.Context, queryName string, query Queries) (sql.Result, error) {
	sqlQuery, args := query.GetSQLQuery(queryName)
	return d.db.ExecContext(ctx, sqlQuery, args...)
}

func (d *Database) Query(ctx context.Context, queryName string, query Queries) (*sql.Rows, error) {
	sqlQuery, args := query.GetSQLQuery(queryName)
	return d.db.QueryContext(ctx, sqlQuery, args...)
}

func (d *Database) QueryRow(ctx context.Context, queryName string, query Queries) *sql.Row {
	sqlQuery, args := query.GetSQLQuery(queryName)
	return d.db.QueryRowContext(ctx, sqlQuery, args...)
}
