-- +goose Up

ALTER TABLE vpn_configs
ADD COLUMN device_id UUID;

CREATE UNIQUE INDEX IF NOT EXISTS vpn_configs_device_protocol_idx
ON vpn_configs(device_id, protocol);

-- +goose Down

DROP INDEX IF EXISTS vpn_configs_device_protocol_idx;

ALTER TABLE vpn_configs
DROP COLUMN device_id;