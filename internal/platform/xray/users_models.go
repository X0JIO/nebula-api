package xray

type AddUserRequest struct {
	UUID  string `json:"uuid"`
	Email string `json:"email"`

	InboundID int `json:"inbound_id"`

	Protocol string `json:"protocol"`

	AlterID int `json:"alter_id,omitempty"`

	Flow string `json:"flow,omitempty"`

	Password string `json:"password,omitempty"`

	Method string `json:"method,omitempty"`
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
