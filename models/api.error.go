package models

import "time"

type ApiError struct {
	Timestamp  time.Time `json:"timestamp"`
	Message    string    `json:"message"`
	StatusCode int       `json:"status_code"`
}
