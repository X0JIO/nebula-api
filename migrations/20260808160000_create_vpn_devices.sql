-- +goose Up

CREATE TABLE vpn_devices (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    vpn_user_id UUID NOT NULL
        REFERENCES vpn_users(id)
        ON DELETE CASCADE,

    name TEXT NOT NULL,

    platform TEXT NOT NULL,

    device_token TEXT NOT NULL UNIQUE,

    last_seen_at TIMESTAMPTZ,

    revoked BOOLEAN NOT NULL DEFAULT FALSE,

    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);


CREATE INDEX idx_vpn_devices_user
ON vpn_devices(vpn_user_id);


-- +goose Down

DROP TABLE vpn_devices;