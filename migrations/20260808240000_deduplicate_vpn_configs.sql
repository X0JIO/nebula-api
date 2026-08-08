-- +goose Up

DELETE FROM vpn_configs a
USING vpn_configs b
WHERE
    a.protocol = b.protocol
    AND a.device_id = b.device_id
    AND a.created_at < b.created_at;


-- +goose Down