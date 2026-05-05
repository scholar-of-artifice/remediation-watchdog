package worker

import (
	"context"
	"eager-marmot/internal/models"
	"eager-marmot/internal/store"
	"encoding/json"
	"log"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type Processor struct {
	Redis *store.RedisStore
}

// ProcessMessage handles the transformation from Kafka raw bytes to the structured Redis sink.
// This serves as a validation gate, ensuring only messages that match the EventRequest schema are persisted to the final data store.
func (p *Processor) ProcessMessage(ctx context.Context, msg *kafka.Message) {
	// unmarshal to validate the contract
	var event models.EventRequest
	// did we get an error?
	if err := json.Unmarshal(msg.Value, &event); err != nil {
		log.Printf("Error decoding message: %v", err)
		return
	}
	// pass the ID and raw JSON to the store
	// save the raw string to Redis
	// you can pretend to use any other database if you want...
	err := p.Redis.SaveEvent(ctx, event.ID, msg.Value)
	// did we get an error?
	if err != nil {
		log.Printf("Failed to persist event %s to Redis: %v", event.ID, err)
		return
	}
	log.Printf("Successfully processed event %s from partition %d", event.ID, msg.TopicPartition.Partition)
}
