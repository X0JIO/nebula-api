package sessions

import (
	db "github.com/X0JIO/nebula-api/internal/platform/database/sqlc"
)

func ToResponse(s db.Session) SessionResponse {

	var device string
	if s.DeviceName.Valid {
		device = s.DeviceName.String
	}

	var agent string
	if s.UserAgent.Valid {
		agent = s.UserAgent.String
	}

	var ip string
	if s.Ip != nil {
		ip = s.Ip.String()
	}

	return SessionResponse{
		ID:         s.ID.String(),
		UserID:     s.UserID.String(),
		DeviceName: device,
		UserAgent:  agent,
		IPAddress:  ip,
		ExpiresAt:  s.ExpiresAt.Time,
		LastSeen:   s.LastSeen.Time,
		Revoked:    s.Revoked,
		CreatedAt:  s.CreatedAt.Time,
	}
}

func ToResponses(list []db.Session) []SessionResponse {

	resp := make([]SessionResponse, 0, len(list))

	for _, s := range list {
		resp = append(resp, ToResponse(s))
	}

	return resp
}
