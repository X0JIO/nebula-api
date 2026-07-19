package tasks

import (
	"context"

	db "github.com/X0JIO/nebula-api/internal/platform/database/sqlc"

	"github.com/jackc/pgx/v5/pgtype"
)

type Repository struct {
	q *db.Queries
}

func NewRepository(q *db.Queries) *Repository {
	return &Repository{
		q: q,
	}
}

func (r *Repository) Create(
	ctx context.Context,
	arg db.CreateTaskParams,
) (db.Task, error) {

	return r.q.CreateTask(
		ctx,
		arg,
	)
}

func (r *Repository) GetByID(
	ctx context.Context,
	id pgtype.UUID,
) (db.Task, error) {

	return r.q.GetTaskByID(
		ctx,
		id,
	)
}

func (r *Repository) ListByProject(
	ctx context.Context,
	projectID pgtype.UUID,
) ([]db.Task, error) {

	return r.q.ListTasksByProject(
		ctx,
		projectID,
	)
}

func (r *Repository) ListByAssignee(
	ctx context.Context,
	userID pgtype.UUID,
) ([]db.Task, error) {

	return r.q.ListTasksByAssignee(
		ctx,
		userID,
	)
}

func (r *Repository) ListByStatus(
	ctx context.Context,
	projectID pgtype.UUID,
	status string,
) ([]db.Task, error) {

	return r.q.ListTasksByStatus(
		ctx,
		db.ListTasksByStatusParams{
			ProjectID: projectID,
			Status:    status,
		},
	)
}

func (r *Repository) Update(
	ctx context.Context,
	arg db.UpdateTaskParams,
) (db.Task, error) {

	return r.q.UpdateTask(
		ctx,
		arg,
	)
}

func (r *Repository) Delete(
	ctx context.Context,
	id pgtype.UUID,
) error {

	return r.q.DeleteTask(
		ctx,
		id,
	)
}
