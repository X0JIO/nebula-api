-- name: GetUserByID :one

SELECT
    id,
    email,
    password_hash,
    status,
    created_at,
    updated_at,
    role
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
    updated_at,
    role
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
    updated_at,
    role
FROM users
ORDER BY created_at DESC;



-- name: CreateUser :one

INSERT INTO users (
    email,
    password_hash,
    role
)
VALUES (
    $1,
    $2,
    'user'
)
RETURNING
    id,
    email,
    password_hash,
    status,
    created_at,
    updated_at,
    role;