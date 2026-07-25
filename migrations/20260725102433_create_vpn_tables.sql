-- +goose Up

CREATE TABLE vpn_users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    user_id UUID NOT NULL UNIQUE
        REFERENCES users(id)
        ON DELETE CASCADE,

    uuid UUID NOT NULL,

    private_key TEXT NOT NULL,

    public_key TEXT NOT NULL,

    short_id VARCHAR(32) NOT NULL,

    subscription_token UUID NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE vpn_configs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    vpn_user_id UUID NOT NULL
        REFERENCES vpn_users(id)
        ON DELETE CASCADE,

    protocol TEXT NOT NULL,

    config TEXT NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),

    UNIQUE(vpn_user_id, protocol)
);

CREATE INDEX idx_vpn_users_user_id
ON vpn_users(user_id);

CREATE INDEX idx_vpn_users_subscription
ON vpn_users(subscription_token);

CREATE INDEX idx_vpn_configs_user
ON vpn_configs(vpn_user_id);

-- +goose Down

DROP TABLE vpn_configs;
DROP TABLE vpn_users;