CREATE OR REPLACE PROCEDURE accounts.usp_create_token(
    p_token text,
    p_user_id uuid,
    p_expires_at timestamp
) 
AS $$
    INSERT INTO accounts.tokens (token, user_id, expires_at) VALUES (p_token, p_user_id, p_expires_at);
$$ LANGUAGE sql;


-- TODO: Update the name to all revoke_all, remove old one
CREATE OR REPLACE PROCEDURE accounts.usp_revoke_token_user(
    p_user_id uuid
) 
AS $$
    UPDATE accounts.tokens SET valid = FALSE WHERE user_id = p_user_id;
$$ LANGUAGE sql;

CREATE OR REPLACE PROCEDURE accounts.usp_revoke_token(
    p_token text
) 
AS $$
    UPDATE accounts.tokens SET valid = FALSE WHERE token = p_token;
$$ LANGUAGE sql;

CREATE OR REPLACE FUNCTION accounts.usf_is_token_valid(
    p_token text
) RETURNS boolean
AS $$
    SELECT valid FROM accounts.tokens WHERE token = p_token AND expires_at > CURRENT_TIMESTAMP AND valid = TRUE;
$$ LANGUAGE sql;

CREATE OR REPLACE FUNCTION accounts.usf_is_token_valid_user(
    p_user_id uuid
) RETURNS boolean
AS $$
    SELECT valid FROM accounts.tokens WHERE user_id = p_user_id AND expires_at > CURRENT_TIMESTAMP AND valid = TRUE;
$$ LANGUAGE sql;

CREATE OR REPLACE FUNCTION accounts.usf_remove_invalid_tokens()
RETURNS integer
AS $$
DECLARE
    v_row_count integer;
BEGIN
    DELETE FROM accounts.tokens 
    WHERE expires_at < CURRENT_TIMESTAMP OR valid = FALSE;
    GET DIAGNOSTICS v_row_count = ROW_COUNT;
    RETURN v_row_count;
END;
$$ LANGUAGE plpgsql;