-- name: CreateDevice :one
INSERT INTO devices (
    user_id,
    session_id,
    name,
    platform,
    fingerprint,
    vpn_uuid,
    last_ip
)
VALUES (
    $1,$2,$3,$4,$5,$6,$7
)
RETURNING *;

-- name: GetDevice :one
SELECT *
FROM devices
WHERE id=$1;

-- name: ListDevices :many
SELECT *
FROM devices
WHERE user_id=$1
ORDER BY last_seen DESC;

-- name: UpdateDeviceLastSeen :exec
UPDATE devices
SET
    last_seen = now(),
    last_ip = $2
WHERE id = $1;

-- name: DeleteDevice :exec
DELETE
FROM devices
WHERE id=$1;

-- name: DeleteUserDevices :exec
DELETE
FROM devices
WHERE user_id=$1;

-- name: GetDeviceByFingerprint :one
SELECT *
FROM devices
WHERE user_id = $1
  AND fingerprint = $2;

-- name: UpdateDeviceSession :exec
UPDATE devices
SET
    session_id = $2,
    last_seen = now(),
    last_ip = $3
WHERE id = $1;

-- name: GetDeviceBySessionID :one
SELECT *
FROM devices
WHERE session_id = $1;