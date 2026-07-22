-- +goose Up

CREATE TABLE devices (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),

    user_id uuid NOT NULL
        REFERENCES users(id)
        ON DELETE CASCADE,

    session_id uuid
        REFERENCES sessions(id)
        ON DELETE SET NULL,

    name text NOT NULL,

    platform text NOT NULL,

    fingerprint text NOT NULL,

    vpn_uuid uuid,

    last_ip inet,

    last_seen timestamptz NOT NULL DEFAULT now(),

    created_at timestamptz NOT NULL DEFAULT now(),

    UNIQUE(user_id,fingerprint)
);

CREATE INDEX idx_devices_user
ON devices(user_id);

-- +goose Down

DROP TABLE devices;