CREATE TABLE vpn_servers (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    name TEXT NOT NULL,

    host TEXT NOT NULL,

    port INTEGER NOT NULL DEFAULT 443,

    country TEXT NOT NULL,

    public_key TEXT,

    private_key TEXT,

    short_id TEXT,

    status TEXT NOT NULL DEFAULT 'active',

    capacity INTEGER NOT NULL DEFAULT 100,

    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);


CREATE INDEX idx_vpn_servers_status
ON vpn_servers(status);