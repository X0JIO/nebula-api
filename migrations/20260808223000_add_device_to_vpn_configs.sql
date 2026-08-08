-- +goose Up

ALTER TABLE vpn_configs
ADD COLUMN IF NOT EXISTS device_id UUID;


ALTER TABLE vpn_configs
DROP CONSTRAINT IF EXISTS vpn_configs_vpn_user_id_protocol_key;


ALTER TABLE vpn_configs
ADD CONSTRAINT vpn_configs_user_device_protocol_unique
UNIQUE (vpn_user_id, device_id, protocol);


ALTER TABLE vpn_configs
ADD CONSTRAINT vpn_configs_device_fk
FOREIGN KEY (device_id)
REFERENCES vpn_devices(id)
ON DELETE CASCADE;


CREATE INDEX IF NOT EXISTS idx_vpn_configs_device
ON vpn_configs(device_id);


-- +goose Down

DROP INDEX IF EXISTS idx_vpn_configs_device;

ALTER TABLE vpn_configs
DROP CONSTRAINT IF EXISTS vpn_configs_user_device_protocol_unique;

ALTER TABLE vpn_configs
DROP CONSTRAINT IF EXISTS vpn_configs_device_fk;

ALTER TABLE vpn_configs
DROP COLUMN IF EXISTS device_id;