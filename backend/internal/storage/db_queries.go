package storage

type Queries interface {
	// params: procName
	//
	// return: SQL query and args
	GetSQLQuery(string) (string, []any)
}

const (
	CreateUser = `
	CALL accounts.upc_create_user(
			$1::varchar,
			$2::varchar
	);
	`

	IsUserExists = `
	SELECT accounts.ufn_is_user_exists($1::varchar);
	`
	/// NOT DONE YET

	GetUserByEmail = `
	SELECT * FROM auth.GetUserByEmail($1::varchar);
	`
	GetUsers = `
	SELECT id, email, password, name, created_at FROM auth.users;
	`
)
