CREATE TABLE IF NOT EXISTS accounts.users_images (
    user_id uuid NOT NULL,
    image_path text NOT NULL UNIQUE,
    FOREIGN KEY (user_id) 
        REFERENCES accounts.users(id)
        ON DELETE CASCADE
);


-- FUNCTIONS

CREATE OR REPLACE PROCEDURE accounts.usp_create_user_image(
    p_user_id uuid,
    p_image_path text
)
AS $$
    INSERT INTO accounts.users_images (user_id, image_path) VALUES (p_user_id, p_image_path);
$$ LANGUAGE sql;

CREATE OR REPLACE FUNCTION accounts.usf_get_user_image(
    p_user_id uuid
) RETURNS text 
AS $$
    SELECT image_path FROM accounts.users_images WHERE user_id = p_user_id;
$$ LANGUAGE sql;