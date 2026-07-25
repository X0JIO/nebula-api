package xray

import "context"

type StubManager struct{}

func NewStubManager() *StubManager {
	return &StubManager{}
}

func (*StubManager) Reload(context.Context) error {
	return nil
}

func (*StubManager) Restart(context.Context) error {
	return nil
}

func (*StubManager) Start(context.Context) error {
	return nil
}

func (*StubManager) Stop(context.Context) error {
	return nil
}

func (*StubManager) Health(context.Context) error {
	return nil
}
