package xray

import "context"

type StubClient struct{}

func NewStubClient() Client {
	return &StubClient{}
}

// --------------------------------------------------
// Health
// --------------------------------------------------

func (s *StubClient) Ping(ctx context.Context) error {
	return nil
}

// --------------------------------------------------
// Users
// --------------------------------------------------

func (s *StubClient) AddUser(ctx context.Context, req AddUserRequest) error {
	return nil
}

func (s *StubClient) UpdateUser(ctx context.Context, req UpdateUserRequest) error {
	return nil
}

func (s *StubClient) DeleteUser(ctx context.Context, uuid string) error {
	return nil
}

func (s *StubClient) GetUserStats(ctx context.Context, uuid string) (*UserStats, error) {
	return &UserStats{}, nil
}

// --------------------------------------------------
// Inbounds
// --------------------------------------------------

func (s *StubClient) ListInbounds(ctx context.Context) ([]Inbound, error) {
	return []Inbound{}, nil
}

func (s *StubClient) FindInboundByProtocol(ctx context.Context, protocol string) (*Inbound, error) {
	return &Inbound{
		ID:       1,
		Tag:      protocol,
		Protocol: protocol,
		Port:     443,
	}, nil
}

func (s *StubClient) CreateInbound(ctx context.Context, req CreateInboundRequest) error {
	return nil
}

func (s *StubClient) UpdateInbound(ctx context.Context, req UpdateInboundRequest) error {
	return nil
}

func (s *StubClient) DeleteInbound(ctx context.Context, id int) error {
	return nil
}
