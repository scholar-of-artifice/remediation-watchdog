package api

import (
	"absurd-iguana/internal/models"
	"absurd-iguana/internal/store"
	"encoding/json"
	"net/http"
)

type Handler struct {
	Repo *store.RedisStore
}

func (h *Handler) ProduceEventHandler(w http.ResponseWriter, r *http.Request) {
	var req models.EventRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	//verify connection
	err := h.Repo.SaveEvent(r.Context(), req.ID, req.Payload)
	if err != nil {
		http.Error(w, "Failed to persist to Redis", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(models.EventResponse{
		Status:  "Success",
		Message: "Event received and persisted to Redis",
	})
}
