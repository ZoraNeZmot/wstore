package dbusers

const (
	CreateUser = `
	SELECT * FROM accounts.usf_create_user(
			$1::varchar,
			$2::varchar,
			$3::varchar
	);
	`

	CreateOAuthUser = `
	SELECT * FROM accounts.usf_create_oauth_user(
			$1::varchar,
			$2::varchar
	);
	`

	IsUserExists = `
	SELECT * FROM accounts.usf_is_user_exists($1::varchar);
	`
	GetUserByEmail = `
	SELECT * FROM accounts.usf_get_user_by_email($1::varchar);
	`
	GetUserByid = `
	SELECT * FROM accounts.usf_get_user_by_id($1::uuid);
	`
	GetUsers = `
	SELECT id, email, name FROM accounts.users;
	`
	GetUserProfileImage = `
	SELECT * FROM accounts.usf_get_user_image($1::uuid);
	`
	CreateUserProfileImage = `
	CALL accounts.usp_create_user_image(
		$1::uuid,
		$2::text
	);
	`
)
