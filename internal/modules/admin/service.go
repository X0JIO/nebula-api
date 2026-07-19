package admin

import (
	"context"
	"errors"

	db "github.com/X0JIO/nebula-api/internal/platform/database/sqlc"
)

var (
	ErrInvalidRole   = errors.New("invalid role")
	ErrInvalidStatus = errors.New("invalid status")
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

func (s *Service) ListUsers(
	ctx context.Context,
) ([]db.User, error) {

	return s.repository.ListUsers(ctx)
}

func (s *Service) GetUser(
	ctx context.Context,
	id string,
) (db.User, error) {

	return s.repository.GetUserByID(
		ctx,
		id,
	)
}

func (s *Service) ChangeRole(
	ctx context.Context,
	id string,
	role string,
) (db.User, error) {

	switch role {

	case "user":
	case "admin":

	default:
		return db.User{}, ErrInvalidRole
	}

	return s.repository.UpdateUserRole(
		ctx,
		id,
		role,
	)
}

func (s *Service) ChangeStatus(
	ctx context.Context,
	id string,
	status string,
) (db.User, error) {

	switch status {

	case "active":
	case "blocked":

	default:
		return db.User{}, ErrInvalidStatus
	}

	return s.repository.UpdateUserStatus(
		ctx,
		id,
		status,
	)
}

func (s *Service) DeleteUser(
	ctx context.Context,
	id string,
) error {

	return s.repository.DeleteUser(
		ctx,
		id,
	)
}

type Dashboard struct {
	Users          int64 `json:"users"`
	Admins         int64 `json:"admins"`
	BlockedUsers   int64 `json:"blocked_users"`
	ActiveSessions int64 `json:"active_sessions"`
}

func (s *Service) Dashboard(
	ctx context.Context,
) (Dashboard, error) {

	stats, err := s.repository.Dashboard(ctx)

	if err != nil {
		return Dashboard{}, err
	}

	return Dashboard{
		Users:          stats.Users,
		Admins:         stats.Admins,
		BlockedUsers:   stats.BlockedUsers,
		ActiveSessions: stats.ActiveSessions,
	}, nil
}
