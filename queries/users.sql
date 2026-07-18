-- name: GetUserByID :one

SELECT
    id,
    email,
    password_hash,
    status,
    created_at,
    updated_at
FROM users
WHERE id = $1
LIMIT 1;



-- name: GetUserByEmail :one

SELECT
    id,
    email,
    password_hash,
    status,
    created_at,
    updated_at
FROM users
WHERE email = $1
LIMIT 1;



-- name: ListUsers :many

SELECT
    id,
    email,
    password_hash,
    status,
    created_at,
    updated_at
FROM users
ORDER BY created_at DESC;



-- name: CreateUser :one

INSERT INTO users (
    email,
    password_hash
)
VALUES (
    $1,
    $2
)
RETURNING
    id,
    email,
    password_hash,
    status,
    created_at,
    updated_at;