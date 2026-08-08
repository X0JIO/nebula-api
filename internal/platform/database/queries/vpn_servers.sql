-- name: CreateVPNServer :one
INSERT INTO vpn_servers (
    name,
    host,
    port,
    country,
    public_key,
    private_key,
    short_id
)
VALUES (
    $1, $2, $3, $4, $5, $6, $7
)
RETURNING *;


-- name: ListVPNServers :many
SELECT *
FROM vpn_servers
ORDER BY created_at DESC;


-- name: GetVPNServer :one
SELECT *
FROM vpn_servers
WHERE id = $1
LIMIT 1;


-- name: GetActiveVPNServer :one
SELECT *
FROM vpn_servers
WHERE status = 'active'
ORDER BY capacity ASC
LIMIT 1;


-- name: UpdateVPNServer :one
UPDATE vpn_servers
SET
    name = $2,
    host = $3,
    port = $4,
    country = $5,
    public_key = $6,
    private_key = $7,
    short_id = $8
WHERE id = $1
RETURNING *;


-- name: DeleteVPNServer :exec
DELETE FROM vpn_servers
WHERE id = $1;


-- name: DeactivateAllVPNServers :exec
UPDATE vpn_servers
SET status = 'inactive'
WHERE status = 'active';


-- name: ActivateVPNServer :one
UPDATE vpn_servers
SET status = 'active'
WHERE id = $1
RETURNING *;