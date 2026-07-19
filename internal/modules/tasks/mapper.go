package tasks

import (
	db "github.com/X0JIO/nebula-api/internal/platform/database/sqlc"

	"github.com/jackc/pgx/v5/pgtype"
)

func ToCreateParams(
	projectID pgtype.UUID,
	creatorID pgtype.UUID,
	req CreateTaskRequest,
) db.CreateTaskParams {

	var assignee pgtype.UUID
	if req.AssigneeID != nil {
		assignee = *req.AssigneeID
	}

	var due pgtype.Timestamptz
	if req.DueDate != nil {
		due = *req.DueDate
	}

	return db.CreateTaskParams{
		ProjectID:   projectID,
		CreatorID:   creatorID,
		Title:       req.Title,
		Description: req.Description,
		AssigneeID:  assignee,
		Status:      StatusTodo,
		Priority:    req.Priority,
		DueDate:     due,
	}
}

func ToUpdateParams(
	id pgtype.UUID,
	req UpdateTaskRequest,
) db.UpdateTaskParams {

	var assignee pgtype.UUID
	if req.AssigneeID != nil {
		assignee = *req.AssigneeID
	}

	var due pgtype.Timestamptz
	if req.DueDate != nil {
		due = *req.DueDate
	}

	return db.UpdateTaskParams{
		ID:          id,
		Title:       req.Title,
		Description: req.Description,
		AssigneeID:  assignee,
		Status:      req.Status,
		Priority:    req.Priority,
		DueDate:     due,
	}
}

func ToResponse(task db.Task) TaskResponse {

	var assignee *pgtype.UUID
	if task.AssigneeID.Valid {
		assignee = &task.AssigneeID
	}

	var due *pgtype.Timestamptz
	if task.DueDate.Valid {
		t := pgtype.Timestamptz{
			Time:  task.DueDate.Time,
			Valid: true,
		}
		due = &t
	}

	return TaskResponse{
		ID:          task.ID,
		ProjectID:   task.ProjectID,
		CreatorID:   task.CreatorID,
		AssigneeID:  assignee,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		Priority:    task.Priority,
		DueDate:     due,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}
}
