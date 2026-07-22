-- name: CreateSession :one
INSERT INTO sessions (
    user_id,
    refresh_token_hash,
    device_name,
    ip,
    user_agent,
    expires_at
)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;


-- name: GetSessionByRefreshHash :one
SELECT *
FROM sessions
WHERE refresh_token_hash = $1
LIMIT 1;


-- name: RevokeSession :exec
UPDATE sessions
SET revoked = true
WHERE id = $1;


-- name: RevokeUserSessions :exec
UPDATE sessions
SET revoked = true
WHERE user_id = $1;


-- name: UpdateSessionRefresh :exec
UPDATE sessions
SET
    refresh_token_hash = $2,
    expires_at = $3,
    last_seen = now()
WHERE id = $1;


-- name: ListSessions :many
SELECT *
FROM sessions
WHERE user_id = $1
AND revoked = false
ORDER BY last_seen DESC;

-- name: DeleteSession :exec
DELETE
FROM sessions
WHERE id = $1;

-- name: GetSession :one
SELECT *
FROM sessions
WHERE id = $1;
