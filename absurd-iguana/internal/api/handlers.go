package api

import (
	"absurd-iguana/internal/models"
	"context"
	"encoding/json"
	"net/http"
)

// EventRepository defines the behaviour required for persisting event data.
// By using an interface, the API layer remains decoupled from the specific storage implementation.
// This allows for the swap with the real KafkaStore for a Mock in tests.
type EventRepository interface {
	SaveEvent(ctx context.Context, id string, data interface{}) error
}

// Handler serves as the primary controller for API interactions.
// It holds a reference to the EventRepository to perform data operations.
type Handler struct {
	Repo EventRepository
}

// ProduceEventHandler processes incoming POST requests to ingest events.
// It handles the full lifecycle of a requet: decoding, validation, persistence via the repository and strucutred response delivery.
func (h *Handler) ProduceEventHandler(w http.ResponseWriter, r *http.Request) {
	var req models.EventRequest
	// try to decode the incoming request
	// did we get an error?
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	// validate against schema
	if err := req.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	// verify connection
	// persist the event using injected repository implementation.
	err := h.Repo.SaveEvent(r.Context(), req.ID, req.Payload)
	// did we get an error?
	if err != nil {
		http.Error(w, "Failed to persist to Kafka", http.StatusInternalServerError)
		return
	}
	// respond with a 202, indicating the event has been successfully handed off
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(models.EventResponse{
		Status:  "Success",
		Message: "Event received and sent to message bus",
	})
}
