package store

import "time"

type Task struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Status    string    `json:"status"`
	CreatedAT time.Time `json:"created_at"`
}
