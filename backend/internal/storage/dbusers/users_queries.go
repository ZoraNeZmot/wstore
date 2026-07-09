package dbusers

const (
	CreateUser = `
	SELECT accounts.ufn_create_user(
			$1::varchar,
			$2::varchar,
			$3::varchar
	);
	`

	IsUserExists = `
	SELECT accounts.ufn_is_user_exists($1::varchar);
	`
	GetUserByEmail = `
	SELECT * FROM accounts.ufn_get_user_by_email($1::varchar);
	`
	/// NOT DONE YET

	GetUsers = `
	SELECT uid, email, name FROM accounts.users;
	`
)
