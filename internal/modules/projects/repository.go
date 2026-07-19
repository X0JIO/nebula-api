package projects

import (
	"context"

	db "github.com/X0JIO/nebula-api/internal/platform/database/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type Repository struct {
	q db.Querier
}

func NewRepository(q db.Querier) *Repository {
	return &Repository{
		q: q,
	}
}

func (r *Repository) CreateProject(
	ctx context.Context,
	arg db.CreateProjectParams,
) (db.Project, error) {

	return r.q.CreateProject(ctx, arg)
}

func (r *Repository) GetProject(
	ctx context.Context,
	id pgtype.UUID,
) (db.Project, error) {

	return r.q.GetProjectByID(ctx, id)
}

func (r *Repository) ListProjects(
	ctx context.Context,
	userID pgtype.UUID,
) ([]db.Project, error) {

	return r.q.ListProjectsByUser(ctx, userID)
}

func (r *Repository) UpdateProject(
	ctx context.Context,
	arg db.UpdateProjectParams,
) (db.Project, error) {

	return r.q.UpdateProject(ctx, arg)
}

func (r *Repository) DeleteProject(
	ctx context.Context,
	id pgtype.UUID,
) error {

	return r.q.DeleteProject(ctx, id)
}

func (r *Repository) AddMember(
	ctx context.Context,
	arg db.AddProjectMemberParams,
) error {

	return r.q.AddProjectMember(ctx, arg)
}

func (r *Repository) RemoveMember(
	ctx context.Context,
	arg db.RemoveProjectMemberParams,
) error {

	return r.q.RemoveProjectMember(ctx, arg)
}

func (r *Repository) GetRole(
	ctx context.Context,
	projectID pgtype.UUID,
	userID pgtype.UUID,
) (string, error) {

	return r.q.GetProjectRole(
		ctx,
		db.GetProjectRoleParams{
			ProjectID: projectID,
			UserID:    userID,
		},
	)
}

func (r *Repository) Exists(
	ctx context.Context,
	projectID pgtype.UUID,
	userID pgtype.UUID,
) (bool, error) {

	return r.q.ProjectExistsForUser(
		ctx,
		db.ProjectExistsForUserParams{
			ProjectID: projectID,
			UserID:    userID,
		},
	)
}
