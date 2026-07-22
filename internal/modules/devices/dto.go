package devices

import "time"

type DeviceResponse struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Platform string    `json:"platform"`
	LastIP   string    `json:"last_ip"`
	LastSeen time.Time `json:"last_seen"`
	Current  bool      `json:"current"`
}
