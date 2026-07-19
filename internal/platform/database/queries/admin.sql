-- name: UpdateUserRole :one
UPDATE users
SET
    role = $2,
    updated_at = now()
WHERE id = $1
RETURNING *;

-- name: UpdateUserStatus :one
UPDATE users
SET
    status = $2,
    updated_at = now()
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE
FROM users
WHERE id = $1;

-- name: DashboardStats :one
SELECT
(
    SELECT COUNT(*) FROM users
) AS users,

(
    SELECT COUNT(*) FROM users
    WHERE role='admin'
) AS admins,

(
    SELECT COUNT(*) FROM users
    WHERE status='blocked'
) AS blocked_users,

(
    SELECT COUNT(*)
    FROM refresh_tokens
    WHERE revoked_at IS NULL
      AND expires_at > now()
) AS active_sessions;