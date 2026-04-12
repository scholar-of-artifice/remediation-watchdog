package models

import "time"

// EventRequest represents the incoming JSON schema from the user/client
type EventRequest struct {
	ID        string    `json:"id"`
	Payload   string    `json:"payload" validate:"required"`
	Timestamp time.Time `json:"timestamp"`
}

// EventResponse is used for consistent API responses
type EventResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
