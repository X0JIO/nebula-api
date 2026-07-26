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
    $1,$2,$3,$4,$5,$6,$7
)
RETURNING *;


-- name: GetActiveVPNServer :one
SELECT *
FROM vpn_servers
WHERE status = 'active'
ORDER BY capacity ASC
LIMIT 1;


-- name: ListVPNServers :many
SELECT *
FROM vpn_servers
ORDER BY created_at DESC;