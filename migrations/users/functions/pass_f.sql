CREATE OR REPLACE PROCEDURE accounts.upc_create_user_pass(
    p_user_id INT,
    p_password TEXT
) 
AS $$
    INSERT INTO accounts.users_pass (user_id, pass) VALUES (p_user_id, p_password);
$$ LANGUAGE sql;

CREATE OR REPLACE FUNCTION accounts.ufn_get_user_pass(
    p_user_id INT
) RETURNS TEXT
AS $$
    SELECT pass FROM accounts.users_pass WHERE user_id = p_user_id;
$$ LANGUAGE sql;