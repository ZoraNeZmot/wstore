package dbtokens

const (
	CreateToken = `
	CALL accounts.usp_create_token(
		$1::text,
		$2::uuid,
		$3::timestamp
	);
	`
	IsTokenValid = `
	SELECT accounts.usf_is_token_valid($1::text);
	`
	RevokeToken = `
	CALL accounts.usp_revoke_token($1::text);
	`
)
