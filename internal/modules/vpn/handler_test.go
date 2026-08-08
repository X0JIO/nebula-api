package vpn

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	db "github.com/X0JIO/nebula-api/internal/platform/database/sqlc"
	"github.com/X0JIO/nebula-api/internal/platform/web/middleware"
	"github.com/jackc/pgx/v5/pgtype"
)

func authContext(ctx context.Context) context.Context {
	return context.WithValue(
		ctx,
		middleware.ContextUserID,
		"11111111-1111-1111-1111-111111111111",
	)
}

func TestHandlerListConfigs(t *testing.T) {

	repo := &mockVPNRepository{
		getVPNUserFn: func(
			ctx context.Context,
			userID string,
		) (db.VpnUser, error) {
			return testVPNUser(), nil
		},
	}

	service := &Service{
		repo: repo,
	}

	handler := NewHandler(service)

	req := httptest.NewRequest(
		http.MethodGet,
		"/vpn/configs",
		nil,
	)

	req = req.WithContext(
		authContext(req.Context()),
	)

	rec := httptest.NewRecorder()

	handler.ListConfigs(
		rec,
		req,
	)

	if rec.Code != http.StatusOK {
		t.Fatalf(
			"expected status 200, got %d",
			rec.Code,
		)
	}

	var body ConfigResponse

	if err := json.NewDecoder(
		rec.Body,
	).Decode(&body); err != nil {
		t.Fatalf(
			"decode response: %v",
			err,
		)
	}
}

func TestHandlerGetAccount(t *testing.T) {

	repo := &mockVPNRepository{
		getVPNUserFn: func(
			ctx context.Context,
			userID string,
		) (db.VpnUser, error) {
			return db.VpnUser{
				ID: pgtype.UUID{
					Bytes: [16]byte{1},
					Valid: true,
				},
				Uuid:              "11111111-1111-1111-1111-111111111111",
				SubscriptionToken: "token",
				PublicKey:         "public",
				ShortID:           "short",
			}, nil
		},
	}

	service := &Service{
		repo: repo,
	}

	handler := NewHandler(service)

	req := httptest.NewRequest(
		http.MethodGet,
		"/vpn/account",
		nil,
	)

	req = req.WithContext(
		authContext(req.Context()),
	)

	rec := httptest.NewRecorder()

	handler.GetAccount(
		rec,
		req,
	)

	if rec.Code != http.StatusOK {
		t.Fatalf(
			"expected status 200, got %d",
			rec.Code,
		)
	}

	if !strings.Contains(
		rec.Body.String(),
		"11111111-1111-1111-1111-111111111111",
	) {
		t.Fatal(
			"response does not contain UUID",
		)
	}
}
