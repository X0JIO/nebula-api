package server

import (
	"context"
	"testing"

	db "github.com/X0JIO/nebula-api/internal/platform/database/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type mockServerRepository struct {
	createFn func(
		context.Context,
		db.CreateVPNServerParams,
	) (db.VpnServer, error)

	listFn func(
		context.Context,
	) ([]db.VpnServer, error)

	getFn func(
		context.Context,
		pgtype.UUID,
	) (db.VpnServer, error)

	getActiveFn func(
		context.Context,
	) (db.VpnServer, error)

	updateFn func(
		context.Context,
		db.UpdateVPNServerParams,
	) (db.VpnServer, error)

	deleteFn func(
		context.Context,
		pgtype.UUID,
	) error

	activateFn func(
		context.Context,
		pgtype.UUID,
	) (db.VpnServer, error)

	deactivateFn func(
		context.Context,
		pgtype.UUID,
	) (db.VpnServer, error)

	deactivateAllFn func(
		context.Context,
	) error
}

func (m *mockServerRepository) Create(
	ctx context.Context,
	params db.CreateVPNServerParams,
) (db.VpnServer, error) {

	if m.createFn != nil {
		return m.createFn(ctx, params)
	}

	return db.VpnServer{}, nil
}

func (m *mockServerRepository) List(
	ctx context.Context,
) ([]db.VpnServer, error) {

	if m.listFn != nil {
		return m.listFn(ctx)
	}

	return nil, nil
}

func (m *mockServerRepository) Get(
	ctx context.Context,
	id pgtype.UUID,
) (db.VpnServer, error) {

	if m.getFn != nil {
		return m.getFn(ctx, id)
	}

	return db.VpnServer{}, nil
}

func (m *mockServerRepository) GetActive(
	ctx context.Context,
) (db.VpnServer, error) {

	if m.getActiveFn != nil {
		return m.getActiveFn(ctx)
	}

	return db.VpnServer{}, nil
}

func (m *mockServerRepository) Update(
	ctx context.Context,
	params db.UpdateVPNServerParams,
) (db.VpnServer, error) {

	if m.updateFn != nil {
		return m.updateFn(ctx, params)
	}

	return db.VpnServer{}, nil
}

func (m *mockServerRepository) Delete(
	ctx context.Context,
	id pgtype.UUID,
) error {

	if m.deleteFn != nil {
		return m.deleteFn(ctx, id)
	}

	return nil
}

func (m *mockServerRepository) Activate(
	ctx context.Context,
	id pgtype.UUID,
) (db.VpnServer, error) {

	if m.activateFn != nil {
		return m.activateFn(ctx, id)
	}

	return db.VpnServer{}, nil
}

func (m *mockServerRepository) Deactivate(
	ctx context.Context,
	id pgtype.UUID,
) (db.VpnServer, error) {

	if m.deactivateFn != nil {
		return m.deactivateFn(ctx, id)
	}

	return db.VpnServer{}, nil
}

func (m *mockServerRepository) DeactivateAll(
	ctx context.Context,
) error {

	if m.deactivateAllFn != nil {
		return m.deactivateAllFn(ctx)
	}

	return nil
}

func TestServiceCreate(t *testing.T) {
	repo := &mockServerRepository{
		createFn: func(
			ctx context.Context,
			params db.CreateVPNServerParams,
		) (db.VpnServer, error) {

			if params.Name != "Germany 1" {
				t.Fatalf(
					"unexpected name: %s",
					params.Name,
				)
			}

			if params.Host != "vpn.example.com" {
				t.Fatalf(
					"unexpected host: %s",
					params.Host,
				)
			}

			if params.Port != 443 {
				t.Fatalf(
					"unexpected port: %d",
					params.Port,
				)
			}

			return db.VpnServer{
				Name:    params.Name,
				Host:    params.Host,
				Port:    params.Port,
				Country: params.Country,
			}, nil
		},
	}

	service := NewService(repo)

	server, err := service.Create(
		context.Background(),
		CreateRequest{
			Name:    "  Germany 1  ",
			Host:    "  vpn.example.com ",
			Port:    443,
			Country: " DE ",
		},
	)

	if err != nil {
		t.Fatalf(
			"Create() error = %v",
			err,
		)
	}

	if server.Name != "Germany 1" {
		t.Fatalf(
			"unexpected server name: %s",
			server.Name,
		)
	}

	if server.Host != "vpn.example.com" {
		t.Fatalf(
			"unexpected server host: %s",
			server.Host,
		)
	}
}

func TestServiceCreateValidation(t *testing.T) {
	repo := &mockServerRepository{
		createFn: func(
			ctx context.Context,
			params db.CreateVPNServerParams,
		) (db.VpnServer, error) {
			t.Fatal("repository Create() should not be called")
			return db.VpnServer{}, nil
		},
	}

	service := NewService(repo)

	tests := []struct {
		name string
		req  CreateRequest
	}{
		{
			name: "empty name",
			req: CreateRequest{
				Host:    "vpn.example.com",
				Port:    443,
				Country: "DE",
			},
		},
		{
			name: "empty host",
			req: CreateRequest{
				Name:    "Germany",
				Port:    443,
				Country: "DE",
			},
		},
		{
			name: "invalid port",
			req: CreateRequest{
				Name:    "Germany",
				Host:    "vpn.example.com",
				Port:    0,
				Country: "DE",
			},
		},
		{
			name: "empty country",
			req: CreateRequest{
				Name: "Germany",
				Host: "vpn.example.com",
				Port: 443,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			_, err := service.Create(
				context.Background(),
				tt.req,
			)

			if err == nil {
				t.Fatal("expected validation error")
			}
		})
	}
}

func TestServiceGet(t *testing.T) {
	expectedID := "11111111-1111-1111-1111-111111111111"

	repo := &mockServerRepository{
		getFn: func(
			ctx context.Context,
			id pgtype.UUID,
		) (db.VpnServer, error) {

			if !id.Valid {
				t.Fatal("uuid should be valid")
			}

			return db.VpnServer{
				ID:   id,
				Name: "Germany",
				Host: "vpn.example.com",
			}, nil
		},
	}

	service := NewService(repo)

	server, err := service.Get(
		context.Background(),
		expectedID,
	)

	if err != nil {
		t.Fatalf(
			"Get() error = %v",
			err,
		)
	}

	if server.Name != "Germany" {
		t.Fatalf(
			"unexpected server name: %s",
			server.Name,
		)
	}

	if server.Host != "vpn.example.com" {
		t.Fatalf(
			"unexpected host: %s",
			server.Host,
		)
	}
}

func TestServiceGetInvalidUUID(t *testing.T) {
	repo := &mockServerRepository{
		getFn: func(
			ctx context.Context,
			id pgtype.UUID,
		) (db.VpnServer, error) {
			t.Fatal("repository Get() should not be called")
			return db.VpnServer{}, nil
		},
	}

	service := NewService(repo)

	_, err := service.Get(
		context.Background(),
		"invalid-id",
	)

	if err == nil {
		t.Fatal("expected invalid UUID error")
	}
}

func TestServiceList(t *testing.T) {
	repo := &mockServerRepository{
		listFn: func(
			ctx context.Context,
		) ([]db.VpnServer, error) {

			return []db.VpnServer{
				{
					Name: "Germany",
					Host: "de.example.com",
				},
				{
					Name: "Netherlands",
					Host: "nl.example.com",
				},
			}, nil
		},
	}

	service := NewService(repo)

	servers, err := service.List(
		context.Background(),
	)

	if err != nil {
		t.Fatalf(
			"List() error = %v",
			err,
		)
	}

	if len(servers) != 2 {
		t.Fatalf(
			"expected 2 servers, got %d",
			len(servers),
		)
	}

	if servers[0].Name != "Germany" {
		t.Fatalf(
			"unexpected first server: %s",
			servers[0].Name,
		)
	}

	if servers[1].Host != "nl.example.com" {
		t.Fatalf(
			"unexpected second host: %s",
			servers[1].Host,
		)
	}
}

func TestServiceUpdate(t *testing.T) {
	id := "11111111-1111-1111-1111-111111111111"

	repo := &mockServerRepository{
		updateFn: func(
			ctx context.Context,
			params db.UpdateVPNServerParams,
		) (db.VpnServer, error) {

			if params.Name != "Germany Updated" {
				t.Fatalf(
					"unexpected name: %s",
					params.Name,
				)
			}

			if params.Host != "de.updated.com" {
				t.Fatalf(
					"unexpected host: %s",
					params.Host,
				)
			}

			if params.Port != 8443 {
				t.Fatalf(
					"unexpected port: %d",
					params.Port,
				)
			}

			if params.Country != "DE" {
				t.Fatalf(
					"unexpected country: %s",
					params.Country,
				)
			}

			return db.VpnServer{
				Name:    params.Name,
				Host:    params.Host,
				Port:    params.Port,
				Country: params.Country,
			}, nil
		},
	}

	service := NewService(repo)

	server, err := service.Update(
		context.Background(),
		id,
		UpdateRequest{
			Name:    "  Germany Updated ",
			Host:    " de.updated.com ",
			Port:    8443,
			Country: " DE ",
		},
	)

	if err != nil {
		t.Fatalf(
			"Update() error = %v",
			err,
		)
	}

	if server.Name != "Germany Updated" {
		t.Fatalf(
			"unexpected response name: %s",
			server.Name,
		)
	}
}

func TestServiceUpdateInvalidUUID(t *testing.T) {
	repo := &mockServerRepository{
		updateFn: func(
			ctx context.Context,
			params db.UpdateVPNServerParams,
		) (db.VpnServer, error) {
			t.Fatal("Update repository should not be called")
			return db.VpnServer{}, nil
		},
	}

	service := NewService(repo)

	_, err := service.Update(
		context.Background(),
		"bad-uuid",
		UpdateRequest{
			Name:    "Germany",
			Host:    "vpn.example.com",
			Port:    443,
			Country: "DE",
		},
	)

	if err == nil {
		t.Fatal("expected invalid UUID error")
	}
}

func TestServiceUpdateValidation(t *testing.T) {
	repo := &mockServerRepository{
		updateFn: func(
			ctx context.Context,
			params db.UpdateVPNServerParams,
		) (db.VpnServer, error) {
			t.Fatal("Update repository should not be called")
			return db.VpnServer{}, nil
		},
	}

	service := NewService(repo)

	_, err := service.Update(
		context.Background(),
		"11111111-1111-1111-1111-111111111111",
		UpdateRequest{
			Name:    "",
			Host:    "vpn.example.com",
			Port:    443,
			Country: "DE",
		},
	)

	if err == nil {
		t.Fatal("expected validation error")
	}
}
