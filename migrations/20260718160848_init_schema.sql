-- +goose Up

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

    email VARCHAR(255) UNIQUE NOT NULL,

    password_hash TEXT NOT NULL,

    status VARCHAR(50) NOT NULL DEFAULT 'active',

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),

    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);


CREATE INDEX idx_users_email
ON users(email);



-- +goose Down

DROP TABLE users;

DROP EXTENSION IF EXISTS "uuid-ossp";