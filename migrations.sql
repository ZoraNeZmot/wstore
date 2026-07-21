\i migrations/db.sql
\c wstore
\i migrations/roles.sql
-- Users table and functions
\i migrations/users/tables/users.sql
\i migrations/users/functions/users_f.sql
-- Password table and functions
\i migrations/users/tables/users_pass.sql
\i migrations/users/functions/pass_f.sql
-- Tokens table and functions
\i migrations/users/tables/tokens.sql
\i migrations/users/functions/tokens_f.sql
-- Users images table and functions
\i migrations/users/tables/users_images.sql
\i migrations/users/functions/images_f.sql
-- Sales table and functions
\i migrations/sales/tables/items.sql
\i migrations/sales/functions/items_f.sql
