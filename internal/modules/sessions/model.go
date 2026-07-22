package sessions

import "time"

type SessionResponse struct {
	ID         string    `json:"id"`
	UserID     string    `json:"user_id"`
	DeviceName string    `json:"device_name"`
	UserAgent  string    `json:"user_agent"`
	IPAddress  string    `json:"ip_address"`
	ExpiresAt  time.Time `json:"expires_at"`
	LastSeen   time.Time `json:"last_seen"`
	Revoked    bool      `json:"revoked"`
	CreatedAt  time.Time `json:"created_at"`
}
