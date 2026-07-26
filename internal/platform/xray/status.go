package xray

type Status struct {
	Running bool `json:"running"`

	PID int `json:"pid"`
}
