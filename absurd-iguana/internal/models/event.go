package models

import (
	"errors"
	"strings"
	"time"
)

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

// Validate checks if the required fields are present and well formed
func (e *EventRequest) Validate() error {
	if strings.TrimSpace(e.ID) == "" {
		return errors.New("event ID is required.")
	}
	if strings.TrimSpace(e.Payload) == "" {
		return errors.New("payload is required")
	}
	return nil
}
