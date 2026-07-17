package dbusers

const (
	CreateUser = `
	SELECT accounts.usf_create_user(
			$1::varchar,
			$2::varchar,
			$3::varchar
	);
	`

	CreateOAuthUser = `
	SELECT accounts.usf_create_oauth_user(
			$1::varchar,
			$2::varchar
	);
	`

	IsUserExists = `
	SELECT accounts.usf_is_user_exists($1::varchar);
	`
	GetUserByEmail = `
	SELECT * FROM accounts.usf_get_user_by_email($1::varchar);
	`
	GetUserByUid = `
	SELECT * FROM accounts.usf_get_user_by_uid($1::varchar);
	`
	GetUsers = `
	SELECT uid, email, name FROM accounts.users;
	`
)
