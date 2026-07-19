package tasks

import (
	"time"

	"github.com/google/uuid"

	db "github.com/X0JIO/nebula-api/internal/platform/database/sqlc"

	"github.com/jackc/pgx/v5/pgtype"
)

func ToCreateParams(
	projectID pgtype.UUID,
	creatorID pgtype.UUID,
	req CreateTaskRequest,
) db.CreateTaskParams {

	var assigneeID pgtype.UUID

	if req.AssigneeID != nil {

		id, err := uuid.Parse(*req.AssigneeID)

		if err == nil {

			assigneeID = pgtype.UUID{
				Bytes: id,
				Valid: true,
			}

		}
	}

	var dueDate pgtype.Timestamptz

	if req.DueDate != nil {

		dueDate = pgtype.Timestamptz{
			Time:  *req.DueDate,
			Valid: true,
		}

	}

	return db.CreateTaskParams{

		ProjectID: projectID,

		CreatorID: creatorID,

		AssigneeID: assigneeID,

		Title: req.Title,

		Description: req.Description,

		Status: StatusTodo,

		Priority: req.Priority,

		DueDate: dueDate,
	}
}

func ToUpdateParams(
	id pgtype.UUID,
	req UpdateTaskRequest,
) db.UpdateTaskParams {

	var assigneeID pgtype.UUID

	if req.AssigneeID != nil {

		value, err := uuid.Parse(*req.AssigneeID)

		if err == nil {

			assigneeID = pgtype.UUID{
				Bytes: value,
				Valid: true,
			}

		}
	}

	var dueDate pgtype.Timestamptz

	if req.DueDate != nil {

		dueDate = pgtype.Timestamptz{
			Time:  *req.DueDate,
			Valid: true,
		}

	}

	return db.UpdateTaskParams{

		ID: id,

		AssigneeID: assigneeID,

		Title: req.Title,

		Description: req.Description,

		Status: req.Status,

		Priority: req.Priority,

		DueDate: dueDate,
	}
}

func ToResponse(
	task db.Task,
) TaskResponse {

	var assigneeID *string

	if task.AssigneeID.Valid {

		value := task.AssigneeID.String()

		assigneeID = &value
	}

	var dueDate *time.Time

	if task.DueDate.Valid {

		value := task.DueDate.Time

		dueDate = &value
	}

	return TaskResponse{

		ID: task.ID.String(),

		ProjectID: task.ProjectID.String(),

		CreatorID: task.CreatorID.String(),

		AssigneeID: assigneeID,

		Title: task.Title,

		Description: task.Description,

		Status: task.Status,

		Priority: task.Priority,

		DueDate: dueDate,

		CreatedAt: task.CreatedAt.Time,

		UpdatedAt: task.UpdatedAt.Time,
	}
}

func ToResponses(
	tasks []db.Task,
) []TaskResponse {

	result := make(
		[]TaskResponse,
		0,
		len(tasks),
	)

	for _, task := range tasks {

		result = append(
			result,
			ToResponse(task),
		)

	}

	return result
}
