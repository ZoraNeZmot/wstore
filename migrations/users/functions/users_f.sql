-- CREATE USER
CREATE OR REPLACE FUNCTION accounts.ufn_create_user(
    p_email text,
    p_password text,
    p_name text
) RETURNS INT
AS $$
DECLARE 
   v_user_id INT;
BEGIN
    INSERT INTO accounts.users (email, name) VALUES (p_email, p_name) RETURNING id INTO v_user_id;
    CALL accounts.upc_create_user_pass(v_user_id, p_password);

    RETURN v_user_id;
END;
$$ LANGUAGE plpgsql;


-- GET USER
CREATE OR REPLACE FUNCTION accounts.ufn_get_user_by_email(
    p_email text
) RETURNS table (id INT, email TEXT, name TEXT, pass TEXT) 
AS $$
DECLARE
    v_user accounts.users;
    v_pass text;
BEGIN
    SELECT * FROM accounts.users WHERE email = p_email INTO v_user;
    v_pass := accounts.ufn_get_user_pass(v_user.id);
    RETURN (v_user.id, v_user.email, v_user.name, v_pass);
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION accounts.ufn_get_user_by_uid(
    p_uid uuid
) RETURNS accounts.users 
AS $$
    SELECT * FROM accounts.users WHERE uid = p_uid;
$$ LANGUAGE sql;


-- CHECK USER EXISTS
CREATE OR REPLACE FUNCTION accounts.ufn_is_user_exists(
    p_email text
) RETURNS boolean
AS $$
    SELECT EXISTS (SELECT 1 FROM accounts.users WHERE email = p_email);
$$ LANGUAGE sql;

CREATE OR REPLACE FUNCTION accounts.ufn_is_user_exists(
    p_uid uuid
) RETURNS boolean
AS $$
    SELECT EXISTS (SELECT 1 FROM accounts.users WHERE uid = p_uid);
$$ LANGUAGE sql;