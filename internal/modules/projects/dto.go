package projects

import (
	db "github.com/X0JIO/nebula-api/internal/platform/database/sqlc"
)

type CreateProjectRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Visibility  string `json:"visibility"`
}

type UpdateProjectRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Visibility  string `json:"visibility"`
}

type AddMemberRequest struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
}

type ProjectResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	OwnerID     string `json:"owner_id"`
	Visibility  string `json:"visibility"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func ToResponse(project db.Project) ProjectResponse {
	return ProjectResponse{
		ID:          project.ID.String(),
		Name:        project.Name,
		Description: project.Description,
		OwnerID:     project.OwnerID.String(),
		Visibility:  project.Visibility,
		CreatedAt:   project.CreatedAt.Time.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:   project.UpdatedAt.Time.Format("2006-01-02T15:04:05Z07:00"),
	}
}

func ToResponses(projects []db.Project) []ProjectResponse {
	result := make([]ProjectResponse, 0, len(projects))

	for _, p := range projects {
		result = append(result, ToResponse(p))
	}

	return result
}
