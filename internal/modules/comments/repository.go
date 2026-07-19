package comments

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

func (r *Repository) Create(
	ctx context.Context,
	arg db.CreateCommentParams,
) (db.Comment, error) {

	return r.queries.CreateComment(
		ctx,
		arg,
	)
}

func (r *Repository) Get(
	ctx context.Context,
	id pgtype.UUID,
) (db.Comment, error) {

	return r.queries.GetComment(
		ctx,
		id,
	)
}

func (r *Repository) ListTask(
	ctx context.Context,
	taskID pgtype.UUID,
) ([]db.Comment, error) {

	return r.queries.ListTaskComments(
		ctx,
		taskID,
	)
}

func (r *Repository) Delete(
	ctx context.Context,
	id pgtype.UUID,
) error {

	return r.queries.DeleteComment(
		ctx,
		id,
	)
}
