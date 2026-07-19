package projects

import (
	"context"
	"errors"

	db "github.com/X0JIO/nebula-api/internal/platform/database/sqlc"

	"github.com/jackc/pgx/v5/pgtype"
)

var (
	ErrInvalidVisibility = errors.New("invalid visibility")
	ErrInvalidRole       = errors.New("invalid role")
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

func validateVisibility(
	visibility string,
) error {

	switch visibility {

	case "private":
	case "team":
	case "public":

	default:
		return ErrInvalidVisibility
	}

	return nil
}

func validateRole(
	role string,
) error {

	switch role {

	case "owner":
	case "admin":
	case "member":
	case "viewer":

	default:
		return ErrInvalidRole
	}

	return nil
}

func (s *Service) CreateProject(
	ctx context.Context,
	name string,
	description string,
	visibility string,
	ownerID string,
) (db.Project, error) {

	if err := validateVisibility(visibility); err != nil {
		return db.Project{}, err
	}

	var uid pgtype.UUID

	if err := uid.Scan(ownerID); err != nil {
		return db.Project{}, err
	}

	project, err := s.repository.CreateProject(
		ctx,
		db.CreateProjectParams{
			Name:        name,
			Description: description,
			OwnerID:     uid,
			Visibility:  visibility,
		},
	)

	if err != nil {
		return db.Project{}, err
	}

	err = s.repository.AddMember(
		ctx,
		db.AddProjectMemberParams{
			ProjectID: project.ID,
			UserID:    uid,
			Role:      "owner",
		},
	)

	if err != nil {
		return db.Project{}, err
	}

	return project, nil
}

func (s *Service) ListProjects(
	ctx context.Context,
	userID string,
) ([]db.Project, error) {

	var uid pgtype.UUID

	if err := uid.Scan(userID); err != nil {
		return nil, err
	}

	return s.repository.ListProjects(
		ctx,
		uid,
	)
}

func (s *Service) GetProject(
	ctx context.Context,
	id string,
) (db.Project, error) {

	var pid pgtype.UUID

	if err := pid.Scan(id); err != nil {
		return db.Project{}, err
	}

	return s.repository.GetProject(
		ctx,
		pid,
	)
}

func (s *Service) UpdateProject(
	ctx context.Context,
	id string,
	name string,
	description string,
	visibility string,
) (db.Project, error) {

	if err := validateVisibility(visibility); err != nil {
		return db.Project{}, err
	}

	var pid pgtype.UUID

	if err := pid.Scan(id); err != nil {
		return db.Project{}, err
	}

	return s.repository.UpdateProject(
		ctx,
		db.UpdateProjectParams{
			ID:          pid,
			Name:        name,
			Description: description,
			Visibility:  visibility,
		},
	)
}

func (s *Service) DeleteProject(
	ctx context.Context,
	id string,
) error {

	var pid pgtype.UUID

	if err := pid.Scan(id); err != nil {
		return err
	}

	return s.repository.DeleteProject(
		ctx,
		pid,
	)
}

func (s *Service) AddMember(
	ctx context.Context,
	projectID string,
	userID string,
	role string,
) error {

	if err := validateRole(role); err != nil {
		return err
	}

	var pid pgtype.UUID
	var uid pgtype.UUID

	if err := pid.Scan(projectID); err != nil {
		return err
	}

	if err := uid.Scan(userID); err != nil {
		return err
	}

	return s.repository.AddMember(
		ctx,
		db.AddProjectMemberParams{
			ProjectID: pid,
			UserID:    uid,
			Role:      role,
		},
	)
}

func (s *Service) RemoveMember(
	ctx context.Context,
	projectID string,
	userID string,
) error {

	var pid pgtype.UUID
	var uid pgtype.UUID

	if err := pid.Scan(projectID); err != nil {
		return err
	}

	if err := uid.Scan(userID); err != nil {
		return err
	}

	return s.repository.RemoveMember(
		ctx,
		db.RemoveProjectMemberParams{
			ProjectID: pid,
			UserID:    uid,
		},
	)
}

func (s *Service) GetProjectRole(
	ctx context.Context,
	userID string,
	projectID string,
) (string, error) {

	var uid pgtype.UUID
	var pid pgtype.UUID

	if err := uid.Scan(userID); err != nil {
		return "", err
	}

	if err := pid.Scan(projectID); err != nil {
		return "", err
	}

	return s.repository.GetRole(
		ctx,
		pid,
		uid,
	)
}

func (s *Service) ProjectExists(
	ctx context.Context,
	userID string,
	projectID string,
) (bool, error) {

	var uid pgtype.UUID
	var pid pgtype.UUID

	if err := uid.Scan(userID); err != nil {
		return false, err
	}

	if err := pid.Scan(projectID); err != nil {
		return false, err
	}

	return s.repository.Exists(
		ctx,
		pid,
		uid,
	)
}
