package xray

import "context"

type HTTPManager struct {
	client Client
}

func NewHTTPManager(client Client) Manager {
	return &HTTPManager{
		client: client,
	}
}

func (m *HTTPManager) Health(ctx context.Context) error {
	return m.client.Ping(ctx)
}

func (m *HTTPManager) Reload(ctx context.Context) error {
	return m.client.Reload(ctx)
}

func (m *HTTPManager) Restart(ctx context.Context) error {
	return m.client.Restart(ctx)
}
