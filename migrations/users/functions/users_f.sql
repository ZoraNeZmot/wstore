-- CREATE USER
CREATE OR REPLACE FUNCTION accounts.usf_create_user(
    p_email text,
    p_password text,
    p_name text
) RETURNS INT
AS $$
DECLARE 
   v_user_id INT;
BEGIN
    INSERT INTO accounts.users (email, name) VALUES (p_email, p_name) RETURNING id INTO v_user_id;
    CALL accounts.uspcreate_user_pass(v_user_id, p_password);

    RETURN v_user_id;
END;
$$ LANGUAGE plpgsql;

-- CREATE OAUTH USER (no password)
CREATE OR REPLACE FUNCTION accounts.usf_create_oauth_user(
    p_email text,
    p_name text
) RETURNS uuid
AS $$
DECLARE
    v_uid uuid;
BEGIN
    INSERT INTO accounts.users (email, name, email_verified)
    VALUES (p_email, p_name, TRUE)
    RETURNING uid INTO v_uid;

    RETURN v_uid;
END;
$$ LANGUAGE plpgsql;


-- GET USER
CREATE OR REPLACE FUNCTION accounts.usf_get_user_by_email(
    p_email text
) RETURNS table (id uuid, email TEXT, pass TEXT, name TEXT) 
AS $$
DECLARE
    v_user accounts.users;
    v_pass text;
BEGIN
    SELECT * FROM accounts.users WHERE users.email = p_email INTO v_user;
    IF NOT FOUND THEN
        RETURN;
    END IF;
    v_pass := COALESCE(accounts.usf_get_user_pass(v_user.id), '');
    RETURN QUERY SELECT v_user.uid, v_user.email, v_pass, v_user.name;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION accounts.usf_get_user_by_uid(
    p_uid uuid
) RETURNS table (id uuid, email TEXT, pass TEXT, name TEXT) 
AS $$
DECLARE
    v_user accounts.users;
    v_pass text;
BEGIN
    SELECT * FROM accounts.users WHERE users.uid = p_uid INTO v_user;
    IF NOT FOUND THEN
        RETURN;
    END IF;
    v_pass := COALESCE(accounts.usf_get_user_pass(v_user.id), '');
    RETURN QUERY SELECT v_user.uid, v_user.email, v_pass, v_user.name;
END;
$$ LANGUAGE plpgsql;


-- CHECK USER EXISTS
CREATE OR REPLACE FUNCTION accounts.usf_is_user_exists(
    p_email text
) RETURNS boolean
AS $$
    SELECT EXISTS (SELECT 1 FROM accounts.users WHERE email = p_email);
$$ LANGUAGE sql;

CREATE OR REPLACE FUNCTION accounts.usf_is_user_exists(
    p_uid uuid
) RETURNS boolean
AS $$
    SELECT EXISTS (SELECT 1 FROM accounts.users WHERE uid = p_uid);
$$ LANGUAGE sql;