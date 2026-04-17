package models

import (
	"encoding/json"
	"testing"
	"time"
)

func TestEventRequest_JSON(t *testing.T) {
	// define a fixed time for consistency in testing
	now := time.Date(2026, time.April, 12, 10, 0, 0, 0, time.UTC)
	// list of tables to test with
	tests := []struct {
		name     string
		input    EventRequest
		expected string
	}{
		{
			name: "valid request marshalling",
			input: EventRequest{
				ID:        "event-123",
				Payload:   "some data goes here prob idk",
				Timestamp: now,
			},
			// matching the json tags defined in the Type
			expected: `{"id":"event-123","payload":"some data goes here prob idk","timestamp":"2026-04-12T10:00:00Z"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// test marshalling (struct -> JSON)
			data, err := json.Marshal(tt.input)
			if err != nil {
				t.Fatalf("failed to marshal EventRequest: %v", err)
			}
			if string(data) != tt.expected {
				t.Errorf("got %s, want %s", string(data), tt.expected)
			}
			// test marshalling (JSON -> struct)
			var result EventRequest
			err = json.Unmarshal([]byte(tt.expected), &result)
			if err != nil {
				t.Fatalf("failed to unmarshal JSON: %v", err)
			}
			if result.ID != tt.input.ID {
				t.Errorf("got ID %s, want %s", result.ID, tt.input.ID)
			}
		})
	}
}
