package admin

import (
	"context"

	db "github.com/X0JIO/nebula-api/internal/platform/database/sqlc"

	"github.com/jackc/pgx/v5/pgtype"
)

type Repository struct {
	queries *db.Queries
}

func NewRepository(
	queries *db.Queries,
) *Repository {

	return &Repository{
		queries: queries,
	}
}

func (r *Repository) ListUsers(
	ctx context.Context,
) ([]db.User, error) {

	return r.queries.ListUsers(ctx)
}

func (r *Repository) GetUserByID(
	ctx context.Context,
	id string,
) (db.User, error) {

	var uid pgtype.UUID

	if err := uid.Scan(id); err != nil {
		return db.User{}, err
	}

	return r.queries.GetUserByID(
		ctx,
		uid,
	)
}

func (r *Repository) UpdateUserRole(
	ctx context.Context,
	id string,
	role string,
) (db.User, error) {

	var uid pgtype.UUID

	if err := uid.Scan(id); err != nil {
		return db.User{}, err
	}

	return r.queries.UpdateUserRole(
		ctx,
		db.UpdateUserRoleParams{
			ID:   uid,
			Role: role,
		},
	)
}

func (r *Repository) UpdateUserStatus(
	ctx context.Context,
	id string,
	status string,
) (db.User, error) {

	var uid pgtype.UUID

	if err := uid.Scan(id); err != nil {
		return db.User{}, err
	}

	return r.queries.UpdateUserStatus(
		ctx,
		db.UpdateUserStatusParams{
			ID:     uid,
			Status: status,
		},
	)
}

func (r *Repository) DeleteUser(
	ctx context.Context,
	id string,
) error {

	var uid pgtype.UUID

	if err := uid.Scan(id); err != nil {
		return err
	}

	return r.queries.DeleteUser(
		ctx,
		uid,
	)
}

func (r *Repository) Dashboard(
	ctx context.Context,
) (db.DashboardStatsRow, error) {

	return r.queries.DashboardStats(ctx)
}
