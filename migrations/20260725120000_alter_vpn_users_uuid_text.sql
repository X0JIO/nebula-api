-- +goose Up

ALTER TABLE vpn_users
ALTER COLUMN uuid TYPE text USING uuid::text;

ALTER TABLE vpn_users
ALTER COLUMN subscription_token TYPE text USING subscription_token::text;


-- +goose Down

ALTER TABLE vpn_users
ALTER COLUMN uuid TYPE uuid USING uuid::uuid;

ALTER TABLE vpn_users
ALTER COLUMN subscription_token TYPE uuid USING subscription_token::uuid;