ALTER TABLE vpn_configs
ADD COLUMN device_id UUID;

ALTER TABLE vpn_configs
ADD CONSTRAINT vpn_configs_device_id_fk
FOREIGN KEY (device_id)
REFERENCES vpn_devices(id)
ON DELETE CASCADE;


DROP INDEX IF EXISTS vpn_configs_vpn_user_id_protocol_key;


ALTER TABLE vpn_configs
ADD CONSTRAINT vpn_configs_device_protocol_unique
UNIQUE(device_id, protocol);