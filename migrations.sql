\i migrations/db.sql
\c wstore
\i migrations/roles.sql
-- Users table and functions
\i migrations/users/tables/users.sql
\i migrations/users/functions/users_f.sql
-- Password table and functions
\i migrations/users/tables/users_pass.sql
\i migrations/users/functions/pass_f.sql