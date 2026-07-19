package tasks

import (
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
