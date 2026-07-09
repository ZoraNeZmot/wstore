SELECT 'CREATE DATABASE wstore'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'wstore');\gexec
ALTER DATABASE wstore SET search_path TO accounts, public;

\c wstore;

CREATE SCHEMA IF NOT EXISTS accounts;
