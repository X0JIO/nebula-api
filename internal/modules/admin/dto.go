package admin

import (
	db "github.com/X0JIO/nebula-api/internal/platform/database/sqlc"
)

type UserResponse struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func ToUserResponse(
	user db.User,
) UserResponse {

	return UserResponse{
		ID:        user.ID.String(),
		Email:     user.Email,
		Role:      user.Role,
		Status:    user.Status,
		CreatedAt: user.CreatedAt.Time.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt: user.UpdatedAt.Time.Format("2006-01-02T15:04:05Z07:00"),
	}
}

func ToUsersResponse(
	users []db.User,
) []UserResponse {

	result := make([]UserResponse, 0, len(users))

	for _, u := range users {
		result = append(result, ToUserResponse(u))
	}

	return result
}
