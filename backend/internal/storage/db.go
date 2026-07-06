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

/*
	EMAIL AUTHENTICATION
*/

func (d *Database) VerifyUser(ctx context.Context, user *User) bool {
	var checked bool
	result := d.QueryRow(ctx, "VerifyUser", user)
	result.Scan(&checked)
	return checked
}

func (d *Database) GetUserByEmail(ctx context.Context, user *User) (*User, error) {
	var userData User
	result := d.QueryRow(ctx, "GetUserByEmail", user)
	err := result.Scan(&userData.ID, &userData.Email, &userData.Password, &userData.Name, &userData.CreatedAt)
	return &userData, err
}

func (d *Database) CheckUser(ctx context.Context, user *User) bool {
	var exists bool
	result := d.QueryRow(ctx, "CheckUser", user)
	result.Scan(&exists)
	return exists
}

func (d *Database) AddUser(ctx context.Context, user *User) (int64, error) {
	_, err := d.ExecQuery(ctx, "AddUser", user)
	if err != nil {
		return 0, err
	}
	return 1, nil
}

func (d *Database) GetUsers(ctx context.Context) ([]User, error) {
	var users []User
	result, err := d.Query(ctx, "GetUsers", &User{})
	if err != nil {
		return nil, err
	}
	defer result.Close()
	for result.Next() {
		var user User
		err = result.Scan(&user.ID, &user.Email, &user.Password, &user.Name, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
