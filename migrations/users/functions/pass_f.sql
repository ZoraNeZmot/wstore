CREATE OR REPLACE PROCEDURE accounts.upc_create_user_pass(
    p_user_id INT,
    p_password TEXT
) 
AS $$
    INSERT INTO accounts.users_pass (user_id, password) VALUES (p_user_id, p_password);
$$ LANGUAGE sql;