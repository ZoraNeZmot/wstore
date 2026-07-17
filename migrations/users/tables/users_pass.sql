CREATE TABLE IF NOT EXISTS accounts.users_pass (
    users_id INT PRIMARY KEY,
    pass text NOT NULL,
    FOREIGN KEY (users_id) 
        REFERENCES accounts.users(id)
        ON DELETE CASCADE
);