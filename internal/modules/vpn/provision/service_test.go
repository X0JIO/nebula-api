package provision

import (
	"context"
	"testing"

	"github.com/X0JIO/nebula-api/internal/platform/xray"
)

type mockXrayClient struct {
	xray.Client

	findProtocol string
	findInbound  *xray.Inbound

	addUserRequest xray.AddUserRequest
	addUserCalled  bool
}

func (m *mockXrayClient) FindInboundByProtocol(
	_ context.Context,
	protocol string,
) (*xray.Inbound, error) {
	m.findProtocol = protocol

	return m.findInbound, nil
}

func (m *mockXrayClient) AddUser(
	_ context.Context,
	req xray.AddUserRequest,
) error {
	m.addUserCalled = true
	m.addUserRequest = req

	return nil
}

func TestServiceAdd(t *testing.T) {
	tests := []struct {
		name     string
		protocol string
	}{
		{
			name:     "vless",
			protocol: "vless",
		},
		{
			name:     "reality",
			protocol: "reality",
		},
		{
			name:     "vmess",
			protocol: "vmess",
		},
		{
			name:     "trojan",
			protocol: "trojan",
		},
		{
			name:     "shadowsocks",
			protocol: "shadowsocks",
		},
	}

	const (
		userUUID = "11111111-1111-1111-1111-111111111111"
		email    = "test-user@nebula.local"
	)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := &mockXrayClient{
				findInbound: &xray.Inbound{
					ID: 42,
				},
			}

			service := NewService(
				mock,
				nil,
				nil,
			)

			err := service.Add(
				context.Background(),
				tt.protocol,
				userUUID,
				email,
			)

			if err != nil {
				t.Fatalf("Add() error = %v", err)
			}

			if mock.findProtocol != tt.protocol {
				t.Fatalf(
					"FindInboundByProtocol() protocol = %q, want %q",
					mock.findProtocol,
					tt.protocol,
				)
			}

			if !mock.addUserCalled {
				t.Fatal("AddUser() was not called")
			}

			if mock.addUserRequest.InboundID != 42 {
				t.Fatalf(
					"AddUser() inbound ID = %d, want %d",
					mock.addUserRequest.InboundID,
					42,
				)
			}

			if mock.addUserRequest.Protocol != tt.protocol {
				t.Fatalf(
					"AddUser() protocol = %q, want %q",
					mock.addUserRequest.Protocol,
					tt.protocol,
				)
			}

			if mock.addUserRequest.UUID != userUUID {
				t.Fatalf(
					"AddUser() UUID = %q, want %q",
					mock.addUserRequest.UUID,
					userUUID,
				)
			}

			if mock.addUserRequest.Email != email {
				t.Fatalf(
					"AddUser() email = %q, want %q",
					mock.addUserRequest.Email,
					email,
				)
			}
		})
	}
}

func TestServiceAddUnsupportedProtocol(t *testing.T) {
	mock := &mockXrayClient{}

	service := NewService(
		mock,
		nil,
		nil,
	)

	err := service.Add(
		context.Background(),
		"unknown",
		"11111111-1111-1111-1111-111111111111",
		"test-user@nebula.local",
	)

	if err != xray.ErrUnsupportedProtocol {
		t.Fatalf(
			"Add() error = %v, want %v",
			err,
			xray.ErrUnsupportedProtocol,
		)
	}

	if mock.findProtocol != "" {
		t.Fatalf(
			"FindInboundByProtocol() was called for unsupported protocol: %q",
			mock.findProtocol,
		)
	}

	if mock.addUserCalled {
		t.Fatal("AddUser() was called for unsupported protocol")
	}
}
