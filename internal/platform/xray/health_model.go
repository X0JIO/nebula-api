package xray

type HealthStatus struct {
	Process bool `json:"process"`

	PID int `json:"pid"`

	Version string `json:"version,omitempty"`
}
