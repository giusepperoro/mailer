create extension if not exists pg_stat_statements;

-- creation of account table
-- Creation of accounts table
CREATE TABLE IF NOT EXISTS accounts (
    client_id BIGINT GENERATED ALWAYS AS IDENTITY
    (START WITH 1234 INCREMENT BY 1123),
    balance int NOT NULL,
    PRIMARY KEY(client_id)
    );

CREATE INDEX IF NOT EXISTS client_id on accounts;

-- Creation of orders table
CREATE TABLE IF NOT EXISTS orders (
    client_id BIGINT NOT NULL,
    amount int,
    create_at timestamp NOT NULL,
    looked_at timestamp,
    status int
    );