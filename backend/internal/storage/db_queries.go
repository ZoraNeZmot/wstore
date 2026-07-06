package storage

type Queries interface {
	// params: procName
	//
	// return: SQL query and args
	GetSQLQuery(string) (string, []any)
}

const (
	AddUser = `
	CALL auth.AddUser(
		$1::uuid,
		$2::varchar,
		$3::varchar,
		$4::varchar,
		$5::timestamp
);
	`
	CheckUser = `
	SELECT auth.CheckUser($1::varchar);
	`
	GetUserByEmail = `
	SELECT * FROM auth.GetUserByEmail($1::varchar);
	`
	GetUsers = `
	SELECT id, email, password, name, created_at FROM auth.users;
	`
)
