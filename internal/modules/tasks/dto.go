package tasks

import (
	db "github.com/X0JIO/nebula-api/internal/platform/database/sqlc"

	"github.com/jackc/pgx/v5/pgtype"
)

type CreateTaskRequest struct {
	Title       string              `json:"title"`
	Description string              `json:"description"`
	AssigneeID  *pgtype.UUID        `json:"assignee_id,omitempty"`
	Priority    string              `json:"priority"`
	DueDate     *pgtype.Timestamptz `json:"due_date,omitempty"`
}

type UpdateTaskRequest struct {
	Title       string              `json:"title"`
	Description string              `json:"description"`
	AssigneeID  *pgtype.UUID        `json:"assignee_id,omitempty"`
	Status      string              `json:"status"`
	Priority    string              `json:"priority"`
	DueDate     *pgtype.Timestamptz `json:"due_date,omitempty"`
}

type TaskResponse struct {
	ID          pgtype.UUID         `json:"id"`
	ProjectID   pgtype.UUID         `json:"project_id"`
	CreatorID   pgtype.UUID         `json:"creator_id"`
	AssigneeID  *pgtype.UUID        `json:"assignee_id,omitempty"`
	Title       string              `json:"title"`
	Description string              `json:"description"`
	Status      string              `json:"status"`
	Priority    string              `json:"priority"`
	DueDate     *pgtype.Timestamptz `json:"due_date,omitempty"`
	CreatedAt   pgtype.Timestamptz  `json:"created_at"`
	UpdatedAt   pgtype.Timestamptz  `json:"updated_at"`
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
		CreatedAt: pgtype.Timestamptz{
			Time:  task.CreatedAt.Time,
			Valid: true,
		},
		UpdatedAt: pgtype.Timestamptz{
			Time:  task.UpdatedAt.Time,
			Valid: true,
		},
	}
}
