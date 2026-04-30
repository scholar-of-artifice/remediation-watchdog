package api

import (
	"absurd-iguana/internal/models"
	"context"
	"encoding/json"
	"net/http"
)

// EventRepository defines the behaviour required by the handler.
// This allows for the swap with the real RedisStore for a Mock in tests.
type EventRepository interface {
	SaveEvent(ctx context.Context, id string, data interface{}) error
}

type Handler struct {
	Repo EventRepository
}

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
	//verify connection
	err := h.Repo.SaveEvent(r.Context(), req.ID, req.Payload)
	// did we get an error?
	if err != nil {
		http.Error(w, "Failed to persist to Redis", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(models.EventResponse{
		Status:  "Success",
		Message: "Event received and sent to message bus",
	})
}
