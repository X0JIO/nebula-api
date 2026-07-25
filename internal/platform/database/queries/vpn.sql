-- name: CreateVPNUser :one
INSERT INTO vpn_users (
    user_id,
    uuid,
    private_key,
    public_key,
    short_id,
    subscription_token
)
VALUES (
    $1,$2,$3,$4,$5,$6
)
RETURNING *;

-- name: GetVPNUserByUserID :one
SELECT *
FROM vpn_users
WHERE user_id=$1;

-- name: GetVPNUserBySubscription :one
SELECT *
FROM vpn_users
WHERE subscription_token=$1;

-- name: SaveVPNConfig :one
INSERT INTO vpn_configs(
    vpn_user_id,
    protocol,
    config
)
VALUES(
    $1,$2,$3
)
ON CONFLICT(vpn_user_id,protocol)
DO UPDATE
SET config=EXCLUDED.config
RETURNING *;

-- name: ListVPNConfigs :many
SELECT *
FROM vpn_configs
WHERE vpn_user_id=$1;

-- name: DeleteVPNConfig :exec

DELETE
FROM vpn_configs
WHERE id=$1;

-- name: GetVPNConfig :one

SELECT *
FROM vpn_configs
WHERE id=$1;