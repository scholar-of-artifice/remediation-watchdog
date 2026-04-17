package api

import (
	"absurd-iguana/internal/models"
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockStore struct {
	SaveFunc func(ctx context.Context, id string, data interface{}) error
}

func (m *MockStore) SaveEvent(ctx context.Context, id string, data interface{}) error {
	return m.SaveFunc(ctx, id, data)
}

func TestProduceEventHandler(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    interface{}
		mockSaveErr    error
		expectedStatus int
	}{
		{
			name: "Success - valid payload",
			requestBody: models.EventRequest{
				ID:      "123",
				Payload: "hello world",
			},
			mockSaveErr:    nil,
			expectedStatus: http.StatusAccepted,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// setup mock
			mock := &MockStore{
				SaveFunc: func(ctx context.Context, id string, data interface{}) error {
					return tt.mockSaveErr
				},
			}
			h := &Handler{Repo: mock}
			// create request
			var body []byte
			if s, ok := tt.requestBody.(string); ok {
				body = []byte(s)
			} else {
				body, _ = json.Marshal(tt.requestBody)
			}
			req := httptest.NewRequest(http.MethodPost, "/events", bytes.NewBuffer(body))
			w := httptest.NewRecorder()
			// execute the request
			h.ProduceEventHandler(w, req)
			// verify
			if w.Code != tt.expectedStatus {
				t.Errorf("got status %d, expected %d", w.Code, tt.expectedStatus)
			}
		})
	}
}
