package projects

import (
	"context"

	"github.com/X0JIO/nebula-api/internal/shared/apperrors"

	"github.com/jackc/pgx/v5/pgtype"
)

type PermissionService struct {
	repository *Repository
}

func NewPermissionService(
	repository *Repository,
) *PermissionService {

	return &PermissionService{
		repository: repository,
	}
}

func (s *PermissionService) CanViewProject(
	ctx context.Context,
	projectID pgtype.UUID,
	userID pgtype.UUID,
) error {

	ok, err := s.repository.Exists(
		ctx,
		projectID,
		userID,
	)
	if err != nil {
		return err
	}

	if !ok {
		return apperrors.ErrForbidden
	}

	return nil
}

func (s *PermissionService) CanManageMembers(
	ctx context.Context,
	projectID pgtype.UUID,
	userID pgtype.UUID,
) error {

	role, err := s.repository.GetRole(
		ctx,
		projectID,
		userID,
	)
	if err != nil {
		return err
	}

	switch role {

	case RoleOwner:
		return nil

	case RoleAdmin:
		return nil

	default:
		return apperrors.ErrForbidden
	}
}

func (s *PermissionService) CanDeleteProject(
	ctx context.Context,
	projectID pgtype.UUID,
	userID pgtype.UUID,
) error {

	role, err := s.repository.GetRole(
		ctx,
		projectID,
		userID,
	)
	if err != nil {
		return err
	}

	if role != RoleOwner {
		return apperrors.ErrForbidden
	}

	return nil
}

func (s *PermissionService) CanUpdateProject(
	ctx context.Context,
	projectID pgtype.UUID,
	userID pgtype.UUID,
) error {

	role, err := s.repository.GetRole(
		ctx,
		projectID,
		userID,
	)
	if err != nil {
		return err
	}

	switch role {

	case RoleOwner:
		return nil

	case RoleAdmin:
		return nil

	default:
		return apperrors.ErrForbidden
	}
}
