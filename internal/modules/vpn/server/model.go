package server

import "time"

type Server struct {
	ID string

	Name string

	Host string

	Port int

	Country string

	PublicKey string

	PrivateKey string

	ShortID string

	Status string

	Capacity int

	CreatedAt time.Time
}
