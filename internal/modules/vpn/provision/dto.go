package provision

import (
	"context"
	"sync/atomic"
)

type ProvisionResult struct {
	UUID string

	Port int

	PrivateKey string
	PublicKey  string

	ShortID string
}

type PortAllocator interface {
	Allocate(ctx context.Context) (int, error)
}

type MemoryPortAllocator struct {
	current atomic.Int32
}

func NewMemoryPortAllocator(start int) *MemoryPortAllocator {
	p := &MemoryPortAllocator{}
	p.current.Store(int32(start))
	return p
}

func (p *MemoryPortAllocator) Allocate(ctx context.Context) (int, error) {
	return int(p.current.Add(1)), nil
}
