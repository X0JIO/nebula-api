-- +goose Up

DELETE FROM vpn_configs
WHERE device_id IS NULL;


-- +goose Down