package redis

import (
	"context"

	goredis "github.com/redis/go-redis/v9"

	"github.com/X0JIO/nebula-api/internal/platform/config"
)

type Client struct {
	Client *goredis.Client
}

func New(ctx context.Context, cfg config.RedisConfig) (*Client, error) {

	client := goredis.NewClient(&goredis.Options{
		Addr:     cfg.Host + ":" + cfg.Port,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}

	return &Client{
		Client: client,
	}, nil
}

func (c *Client) Close() error {
	return c.Client.Close()
}
