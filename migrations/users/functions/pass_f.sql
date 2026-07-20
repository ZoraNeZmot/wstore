CREATE OR REPLACE PROCEDURE accounts.usp_create_user_pass(
    p_users_id uuid,
    p_password TEXT
) 
AS $$
    INSERT INTO accounts.users_pass (users_id, pass) VALUES (p_users_id, p_password);
$$ LANGUAGE sql;

CREATE OR REPLACE FUNCTION accounts.usf_get_user_pass(
    p_users_id uuid
) RETURNS TEXT
AS $$
    SELECT pass FROM accounts.users_pass WHERE users_id = p_users_id;
$$ LANGUAGE sql;