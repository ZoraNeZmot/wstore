-- CREATE USER
CREATE OR REPLACE FUNCTION accounts.usf_create_user(
    p_email text,
    p_password text,
    p_name text
) RETURNS uuid
AS $$
DECLARE 
   v_user_id uuid;
BEGIN
    INSERT INTO accounts.users (email, name) VALUES (p_email, p_name) RETURNING id INTO v_user_id;
    CALL accounts.usp_create_user_pass(v_user_id, p_password);

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
    v_id uuid;
BEGIN
    INSERT INTO accounts.users (email, name, email_verified)
    VALUES (p_email, p_name, TRUE)
    RETURNING id INTO v_id;

    RETURN v_id;
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
    RETURN QUERY SELECT v_user.id, v_user.email, v_pass, v_user.name;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION accounts.usf_get_user_by_id(
    p_id uuid
) RETURNS table (id uuid, email TEXT, pass TEXT, name TEXT) 
AS $$
DECLARE
    v_user accounts.users;
    v_pass text;
BEGIN
    SELECT * FROM accounts.users WHERE users.id = p_id INTO v_user;
    IF NOT FOUND THEN
        RETURN;
    END IF;
    v_pass := COALESCE(accounts.usf_get_user_pass(v_user.id), '');
    RETURN QUERY SELECT v_user.id, v_user.email, v_pass, v_user.name;
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
    p_id uuid
) RETURNS boolean
AS $$
    SELECT EXISTS (SELECT 1 FROM accounts.users WHERE id = p_id);
$$ LANGUAGE sql;