package xray

import "context"

type Manager interface {
	Health(ctx context.Context) error

	Reload(ctx context.Context) error
	Restart(ctx context.Context) error

	Version(ctx context.Context) (string, error)
}
