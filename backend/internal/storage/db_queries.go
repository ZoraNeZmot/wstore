package storage

type Queries interface {
	GetSQLQuery() string
	GetArgs() []any
}

const (
	AddUser = `
	INSERT INTO auth.users (uuid, email, password_hash, name, created_at)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING uuid
	`
)
