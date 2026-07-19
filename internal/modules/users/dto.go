package users

import (
	db "github.com/X0JIO/nebula-api/internal/platform/database/sqlc"
)

type UserResponse struct {
	ID        string `json:"id" example:"63dcbc7c-2507-40e6-b771-61621bc6902d"`
	Email     string `json:"email" example:"user@example.com"`
	Status    string `json:"status" example:"active"`
	Role      string `json:"role" example:"user"`
	CreatedAt string `json:"created_at" example:"2026-07-19T10:25:12Z"`
	UpdatedAt string `json:"updated_at" example:"2026-07-19T10:25:12Z"`
}

func ToResponse(user db.User) UserResponse {
	var createdAt string
	if user.CreatedAt.Valid {
		createdAt = user.CreatedAt.Time.Format("2006-01-02T15:04:05Z07:00")
	}

	var updatedAt string
	if user.UpdatedAt.Valid {
		updatedAt = user.UpdatedAt.Time.Format("2006-01-02T15:04:05Z07:00")
	}

	return UserResponse{
		ID:        user.ID.String(),
		Email:     user.Email,
		Status:    user.Status,
		Role:      user.Role,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

func ToResponseSlice(users []db.User) []UserResponse {
	result := make([]UserResponse, 0, len(users))

	for _, user := range users {
		result = append(result, ToResponse(user))
	}

	return result
}
