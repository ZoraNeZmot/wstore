CREATE TABLE IF NOT EXISTS accounts.users_images (
    user_id uuid NOT NULL,
    image_path text NOT NULL UNIQUE,
    FOREIGN KEY (user_id) 
        REFERENCES accounts.users(id)
        ON DELETE CASCADE
)