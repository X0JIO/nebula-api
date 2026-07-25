package xray

import (
	"context"
)

type Controller interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
	Restart(ctx context.Context) error
	Reload(ctx context.Context) error
	Validate(ctx context.Context) error
	Version(ctx context.Context) (string, error)
}
