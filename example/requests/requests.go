package requests

import "time"

type GetUser struct {
	ID string `param:"id"`
}

type EditUser struct {
	ID        string    `param:"id"`
	Name      string    `json:"name"`
	Timestamp time.Time `json:"timestamp"`
}
