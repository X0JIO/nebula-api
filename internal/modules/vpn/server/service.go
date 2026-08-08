package server

import (
	"context"
	"fmt"
	"strings"

	db "github.com/X0JIO/nebula-api/internal/platform/database/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type Service struct {
	repo *Repository
}

func NewService(
	repo *Repository,
) *Service {
	return &Service{
		repo: repo,
	}
}

type CreateRequest struct {
	Name       string
	Host       string
	Port       int
	Country    string
	PublicKey  string
	PrivateKey string
	ShortID    string
}

func (s *Service) Create(
	ctx context.Context,
	req CreateRequest,
) (db.VpnServer, error) {

	req.Name = strings.TrimSpace(req.Name)
	req.Host = strings.TrimSpace(req.Host)
	req.Country = strings.TrimSpace(req.Country)
	req.PublicKey = strings.TrimSpace(req.PublicKey)
	req.PrivateKey = strings.TrimSpace(req.PrivateKey)
	req.ShortID = strings.TrimSpace(req.ShortID)

	if req.Name == "" {
		return db.VpnServer{}, fmt.Errorf("server name is required")
	}

	if req.Host == "" {
		return db.VpnServer{}, fmt.Errorf("server host is required")
	}

	if req.Port <= 0 || req.Port > 65535 {
		return db.VpnServer{}, fmt.Errorf("invalid server port")
	}

	if req.Country == "" {
		return db.VpnServer{}, fmt.Errorf("server country is required")
	}

	return s.repo.Create(
		ctx,
		db.CreateVPNServerParams{
			Name:       req.Name,
			Host:       req.Host,
			Port:       int32(req.Port),
			Country:    req.Country,
			PublicKey:  textValue(req.PublicKey),
			PrivateKey: textValue(req.PrivateKey),
			ShortID:    textValue(req.ShortID),
		},
	)
}

func (s *Service) List(
	ctx context.Context,
) ([]db.VpnServer, error) {

	return s.repo.List(ctx)
}

func textValue(value string) pgtype.Text {
	if value == "" {
		return pgtype.Text{}
	}

	return pgtype.Text{
		String: value,
		Valid:  true,
	}
}
