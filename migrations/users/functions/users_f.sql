-- CREATE USER
CREATE OR REPLACE PROCEDURE accounts.upc_create_user(
    p_uid uuid,
    p_email text,
    p_password text,
    p_name text
) 
AS $$
BEGIN
    INSERT INTO accounts.users (uid, email, password, name) VALUES (p_uid, p_email, p_password, p_name);
END;
$$ LANGUAGE plpgsql;


-- GET USER
CREATE OR REPLACE FUNCTION accounts.ufn_get_user_by_email(
    p_email text
) RETURNS accounts.users 
AS $$
    SELECT * FROM accounts.users WHERE email = p_email;
$$ LANGUAGE sql;

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