package models

import (
	"errors"
	"strings"
	"time"
)

// EventRequest represents the incoming JSON schema from the user/client.
// It uses standard JSON tags for deserialization and includes a custom
// validation tag to signify requirements for the message payload.
type EventRequest struct {
	ID        string    `json:"id"`
	Payload   string    `json:"payload" validate:"required"`
	Timestamp time.Time `json:"timestamp"`
}

// Validate performs a defensive check on the EventRequest data
// before it enters the business logic or persistence layer.
// It ensures that essential identifiers and payloads are not empty
// or blank.
func (e *EventRequest) Validate() error {
	if strings.TrimSpace(e.ID) == "" {
		return errors.New("event ID is required.")
	}
	if strings.TrimSpace(e.Payload) == "" {
		return errors.New("payload is required")
	}
	return nil
}
