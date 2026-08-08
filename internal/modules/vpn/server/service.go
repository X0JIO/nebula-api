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

type UpdateRequest struct {
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

	if err := validateServer(
		req.Name,
		req.Host,
		req.Port,
		req.Country,
	); err != nil {
		return db.VpnServer{}, err
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

func (s *Service) Get(
	ctx context.Context,
	id string,
) (db.VpnServer, error) {

	uuid, err := parseUUID(id)
	if err != nil {
		return db.VpnServer{}, err
	}

	return s.repo.Get(ctx, uuid)
}

func (s *Service) Update(
	ctx context.Context,
	id string,
	req UpdateRequest,
) (db.VpnServer, error) {

	uuid, err := parseUUID(id)
	if err != nil {
		return db.VpnServer{}, err
	}

	req.Name = strings.TrimSpace(req.Name)
	req.Host = strings.TrimSpace(req.Host)
	req.Country = strings.TrimSpace(req.Country)
	req.PublicKey = strings.TrimSpace(req.PublicKey)
	req.PrivateKey = strings.TrimSpace(req.PrivateKey)
	req.ShortID = strings.TrimSpace(req.ShortID)

	if err := validateServer(
		req.Name,
		req.Host,
		req.Port,
		req.Country,
	); err != nil {
		return db.VpnServer{}, err
	}

	return s.repo.Update(
		ctx,
		db.UpdateVPNServerParams{
			ID:         uuid,
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

func (s *Service) Delete(
	ctx context.Context,
	id string,
) error {

	uuid, err := parseUUID(id)
	if err != nil {
		return err
	}

	return s.repo.Delete(ctx, uuid)
}

func (s *Service) Activate(
	ctx context.Context,
	id string,
) (db.VpnServer, error) {

	uuid, err := parseUUID(id)
	if err != nil {
		return db.VpnServer{}, err
	}

	return s.repo.Activate(ctx, uuid)
}

func (s *Service) Deactivate(
	ctx context.Context,
	id string,
) (db.VpnServer, error) {

	uuid, err := parseUUID(id)
	if err != nil {
		return db.VpnServer{}, err
	}

	return s.repo.Deactivate(ctx, uuid)
}

func (s *Service) DeactivateAll(
	ctx context.Context,
) error {

	return s.repo.DeactivateAll(ctx)
}

func validateServer(
	name string,
	host string,
	port int,
	country string,
) error {

	if name == "" {
		return fmt.Errorf("server name is required")
	}

	if host == "" {
		return fmt.Errorf("server host is required")
	}

	if port <= 0 || port > 65535 {
		return fmt.Errorf("invalid server port")
	}

	if country == "" {
		return fmt.Errorf("server country is required")
	}

	return nil
}

func parseUUID(
	id string,
) (pgtype.UUID, error) {

	var uuid pgtype.UUID

	if err := uuid.Scan(id); err != nil {
		return pgtype.UUID{}, fmt.Errorf("invalid server id")
	}

	return uuid, nil
}

func textValue(
	value string,
) pgtype.Text {

	if value == "" {
		return pgtype.Text{}
	}

	return pgtype.Text{
		String: value,
		Valid:  true,
	}
}

func (s *Service) GetActive(
	ctx context.Context,
) (db.VpnServer, error) {
	return s.repo.GetActive(ctx)
}
