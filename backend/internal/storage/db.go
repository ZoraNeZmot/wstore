package storage

import (
	"context"
	"database/sql"
)

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

func (d *Database) ExecQuery(ctx context.Context, query Queries) (sql.Result, error) {
	args := query.GetArgs()
	return d.db.ExecContext(ctx, query.GetSQLQuery(), args...)
}

func (d *Database) AddUser(ctx context.Context, user *User) (int64, error) {
	result, err := d.ExecQuery(ctx, user)
	if err != nil {
	}
	return result.RowsAffected()
}
