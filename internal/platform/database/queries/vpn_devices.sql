-- name: CreateVPNDevice :one

INSERT INTO vpn_devices(
    vpn_user_id,
    name,
    platform,
    device_token
)
VALUES(
    $1,$2,$3,$4
)
RETURNING *;


-- name: ListVPNDevices :many

SELECT *
FROM vpn_devices
WHERE vpn_user_id=$1
ORDER BY created_at DESC;


-- name: RevokeVPNDevice :exec

UPDATE vpn_devices
SET revoked=true
WHERE id=$1;


-- name: DeleteVPNDevice :exec

DELETE
FROM vpn_devices
WHERE id=$1;