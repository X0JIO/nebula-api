package auth

import db "github.com/X0JIO/nebula-api/internal/platform/database/sqlc"

func ToUserResponse(user db.User) UserResponse {
	return UserResponse{
		ID:     user.ID.String(),
		Email:  user.Email,
		Role:   user.Role,
		Status: user.Status,
	}
}

func ToLoginResponse(user db.User, tokens Tokens) LoginResponse {
	return LoginResponse{
		User:         ToUserResponse(user),
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}
}
