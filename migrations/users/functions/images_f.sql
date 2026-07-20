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