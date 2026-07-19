-- +goose Up

ALTER TABLE users
ADD COLUMN role TEXT NOT NULL DEFAULT 'user';

UPDATE users
SET role = 'admin'
WHERE email = 'admin@nebula.com';

CREATE INDEX idx_users_role
ON users(role);

-- +goose Down

DROP INDEX IF EXISTS idx_users_role;

ALTER TABLE users
DROP COLUMN role;