CREATE OR REPLACE PROCEDURE accounts.upc_create_user_pass(
    p_users_id INT,
    p_password TEXT
) 
AS $$
    INSERT INTO accounts.users_pass (users_id, pass) VALUES (p_users_id, p_password);
$$ LANGUAGE sql;

CREATE OR REPLACE FUNCTION accounts.ufn_get_user_pass(
    p_users_id INT
) RETURNS TEXT
AS $$
    SELECT pass FROM accounts.users_pass WHERE users_id = p_users_id;
$$ LANGUAGE sql;