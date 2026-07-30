\i migrations/db.sql
\c wstore
\i migrations/roles.sql


-- Users table and functions
\i migrations/accounts/users.sql
-- Password table and functions
\i migrations/accounts/users_pass.sql
-- Tokens table and functions
\i migrations/accounts/tokens.sql
-- Users images table and functions
\i migrations/accounts/users_images.sql


-- Sales table and functions
\i migrations/sales/items.sql
-- Items images table and functions
\i migrations/sales/items_images.sql