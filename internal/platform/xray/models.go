package xray

type AddUserRequest struct {
	UUID      string `json:"uuid"`
	Email     string `json:"email,omitempty"`
	Protocol  string `json:"protocol"`
	Inbound   string `json:"inbound"`
	ExpiresAt int64  `json:"expires_at,omitempty"`
}

type UpdateUserRequest struct {
	UUID      string `json:"uuid"`
	ExpiresAt int64  `json:"expires_at,omitempty"`
	Enabled   bool   `json:"enabled"`
}

type AddUserResponse struct {
	Success bool `json:"success"`
}

type UpdateUserResponse struct {
	Success bool `json:"success"`
}

type DeleteUserResponse struct {
	Success bool `json:"success"`
}

type UserStats struct {
	Upload   int64 `json:"upload"`
	Download int64 `json:"download"`
	Total    int64 `json:"total"`
}

func (s UserStats) Traffic() int64 {
	return s.Upload + s.Download
}
