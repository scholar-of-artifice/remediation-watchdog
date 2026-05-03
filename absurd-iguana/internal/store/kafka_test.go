package store

import (
	"context"
	"os"
	"testing"
)

func TestKafkaStore_SaveEvent(t *testing.T) {
	// setup: connect to the local Kafka container
	addr := os.Getenv("KAFKA_BOOTSTRAP_SERVERS")
	if addr == "" {
		// fallback for local development
		addr = "localhost:9092"
	}
	topic := "integration-test-topic"
	//
	store, err := NewKafkaStore(addr, topic)
	if err != nil {
		t.Fatalf("Failed to create KafkaStore: %v", err)
	}
	defer store.Close()
	//
	ctx := context.Background()
	// define test data
	testID := "test-msg-123"
	testData := map[string]string{"foo_event": "bar_test"}
	// execute the test
	t.Run("successfully is accepted into event stream", func(t *testing.T) {
		err := store.SaveEvent(ctx, testID, testData)
		if err != nil {
			t.Fatalf("Failed to save event to Kafka: %v", err)
		}
		// synchronize: wait up to 2 seconds for broker to acknowledge the message
		// verification
		unread := store.Producer.Flush(2 * 1000)
		if unread > 0 {
			t.Errorf("Failed to flush all messages:\t %d remaining", unread)
		}
	})
}
