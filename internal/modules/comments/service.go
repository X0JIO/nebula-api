package comments

import (
	"context"

	db "github.com/X0JIO/nebula-api/internal/platform/database/sqlc"
	"github.com/X0JIO/nebula-api/internal/shared/apperrors"

	"github.com/jackc/pgx/v5/pgtype"
)

type Service struct {
	repository *Repository
}

func NewService(
	repository *Repository,
) *Service {

	return &Service{
		repository: repository,
	}
}

func (s *Service) Create(
	ctx context.Context,
	arg db.CreateCommentParams,
) (db.Comment, error) {

	if arg.Body == "" {
		return db.Comment{}, apperrors.ErrTitleRequired
	}

	return s.repository.Create(
		ctx,
		arg,
	)
}

func (s *Service) Get(
	ctx context.Context,
	id pgtype.UUID,
) (db.Comment, error) {

	return s.repository.Get(
		ctx,
		id,
	)
}

func (s *Service) ListTask(
	ctx context.Context,
	taskID pgtype.UUID,
) ([]db.Comment, error) {

	return s.repository.ListTask(
		ctx,
		taskID,
	)
}

func (s *Service) Delete(
	ctx context.Context,
	id pgtype.UUID,
) error {

	return s.repository.Delete(
		ctx,
		id,
	)
}
