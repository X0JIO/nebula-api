package tasks

import "time"

type CreateTaskRequest struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	AssigneeID  *string    `json:"assignee_id,omitempty"`
	Priority    string     `json:"priority"`
	DueDate     *time.Time `json:"due_date,omitempty"`
}

type UpdateTaskRequest struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	AssigneeID  *string    `json:"assignee_id,omitempty"`
	Status      string     `json:"status"`
	Priority    string     `json:"priority"`
	DueDate     *time.Time `json:"due_date,omitempty"`
}

type TaskResponse struct {
	ID         string  `json:"id"`
	ProjectID  string  `json:"project_id"`
	CreatorID  string  `json:"creator_id"`
	AssigneeID *string `json:"assignee_id,omitempty"`

	Title       string `json:"title"`
	Description string `json:"description"`

	Status   string `json:"status"`
	Priority string `json:"priority"`

	DueDate *time.Time `json:"due_date,omitempty"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
