-- +goose Up

CREATE TABLE sessions (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),

    user_id uuid NOT NULL
        REFERENCES users(id)
        ON DELETE CASCADE,

    refresh_token_hash text NOT NULL UNIQUE,

    device_name text,

    ip inet,

    user_agent text,

    expires_at timestamptz NOT NULL,

    last_seen timestamptz NOT NULL DEFAULT now(),

    revoked boolean NOT NULL DEFAULT false,

    created_at timestamptz NOT NULL DEFAULT now()
);

CREATE INDEX idx_sessions_user
ON sessions(user_id);

CREATE INDEX idx_sessions_last_seen
ON sessions(last_seen);

-- +goose Down

DROP TABLE sessions;