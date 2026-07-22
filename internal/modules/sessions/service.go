package sessions

import (
	"context"

	db "github.com/X0JIO/nebula-api/internal/platform/database/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) List(
	ctx context.Context,
	userID string,
) ([]db.Session, error) {

	var uid pgtype.UUID

	if err := uid.Scan(userID); err != nil {
		return nil, err
	}

	return s.repository.List(ctx, uid)
}

func (s *Service) Revoke(
	ctx context.Context,
	sessionID string,
) error {

	var sid pgtype.UUID

	if err := sid.Scan(sessionID); err != nil {
		return err
	}

	return s.repository.Revoke(ctx, sid)
}

func (s *Service) RevokeAll(
	ctx context.Context,
	userID string,
) error {

	var uid pgtype.UUID

	if err := uid.Scan(userID); err != nil {
		return err
	}

	return s.repository.RevokeAll(ctx, uid)
}

func (s *Service) Delete(
	ctx context.Context,
	sessionID pgtype.UUID,
) error {

	return s.repository.Delete(
		ctx,
		sessionID,
	)
}
