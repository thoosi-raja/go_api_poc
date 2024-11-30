package models

import "time"

type Task struct {
	Id          int       `json:"id"` // This will allow the Id to be nil if not provided
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedOn   time.Time `json:created_on`
}
