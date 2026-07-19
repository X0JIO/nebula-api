package tasks

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
	arg db.CreateTaskParams,
) (db.Task, error) {

	if arg.Title == "" {
		return db.Task{}, apperrors.ErrTitleRequired
	}

	if err := ValidateStatus(arg.Status); err != nil {
		return db.Task{}, err
	}

	if err := ValidatePriority(arg.Priority); err != nil {
		return db.Task{}, err
	}

	return s.repository.Create(ctx, arg)
}

func (s *Service) Get(
	ctx context.Context,
	id pgtype.UUID,
) (db.Task, error) {

	return s.repository.GetByID(
		ctx,
		id,
	)
}

func (s *Service) ListProject(
	ctx context.Context,
	projectID pgtype.UUID,
) ([]db.Task, error) {

	return s.repository.ListByProject(
		ctx,
		projectID,
	)
}

func (s *Service) ListAssignee(
	ctx context.Context,
	userID pgtype.UUID,
) ([]db.Task, error) {

	return s.repository.ListByAssignee(
		ctx,
		userID,
	)
}

func (s *Service) ListStatus(
	ctx context.Context,
	projectID pgtype.UUID,
	status string,
) ([]db.Task, error) {

	return s.repository.ListByStatus(
		ctx,
		projectID,
		status,
	)
}

func (s *Service) Update(
	ctx context.Context,
	arg db.UpdateTaskParams,
) (db.Task, error) {

	if arg.Title == "" {
		return db.Task{}, apperrors.ErrTitleRequired
	}

	if err := ValidateStatus(arg.Status); err != nil {
		return db.Task{}, err
	}

	if err := ValidatePriority(arg.Priority); err != nil {
		return db.Task{}, err
	}

	return s.repository.Update(ctx, arg)
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
