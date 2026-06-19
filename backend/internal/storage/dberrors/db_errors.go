package dberrors

import (
	"database/sql"
	"errors"
)

var (
	UserAlreadyExists = errors.New("user already exists")
	UserNotFound      = sql.ErrNoRows
)
