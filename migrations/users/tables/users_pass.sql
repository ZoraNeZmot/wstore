CREATE TABLE IF NOT EXISTS accounts.users_pass (
    user_id INT PRIMARY KEY,
    password text NOT NULL,
    FOREIGN KEY (user_id) 
        REFERENCES accounts.users(id)
        ON DELETE CASCADE
);