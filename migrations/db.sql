SELECT 'CREATE DATABASE wstore'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'wstore');\gexec
ALTER DATABASE wstore SET search_path TO accounts, sales, public;

\c wstore;

CREATE SCHEMA IF NOT EXISTS accounts;

CREATE SCHEMA IF NOT EXISTS sales;
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'item_status') THEN
        CREATE TYPE sales.item_status AS ENUM ('draft', 'active', 'inactive', 'sold', 'deleted');
    END IF;
END
$$;
