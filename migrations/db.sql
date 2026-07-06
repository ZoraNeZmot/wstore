SELECT 'CREATE DATABASE wstore'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'wstore');\gexec

\c wstore;

CREATE SCHEMA IF NOT EXISTS accounts;
