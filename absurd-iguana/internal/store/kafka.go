package store

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

// KafkaStore manages the publishing of events to a Kafka cluster.
// It implements the store interface required by (this) absurd-iguana service.
type KafkaStore struct {
	Producer *kafka.Producer // the client used to send messages synchronously
	Topic    string          // the destination Kafka topic for events
	// other fields potentially later? ...
}

// NewKafkaStore initializes a new Kafka-backed storage implementation for event remediation.
// It creates a synchronous producer and starts a background worker to handle event processing.
func NewKafkaStore(addr string, topic string) (*KafkaStore, error) {
	// initialize the Kafka producer with core configurations
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": addr,
		"client.id":         "absurd-iguana-producer",
		"acks":              "all",
	})
	// do we have an error?
	if err != nil {
		return nil, fmt.Errorf("failed to create producer: %w", err)
	}
	// assember the store instance
	store := &KafkaStore{
		Producer: p,
		Topic:    topic,
	}
	// start a background process for event management
	go store.executeEventLoop()
	return store, nil
}

// SaveEvent serializes the provided data as JSON and publishes it to the
// configured Kafka topic using the event ID as the message key.
func (s *KafkaStore) SaveEvent(ctx context.Context, id string, data interface{}) error {
	// prepare the message payload
	payload, err := json.Marshal(data)
	// did we get an error?
	if err != nil {
		return fmt.Errorf("failed to marshal event: %w", err)
	}
	// create the kafka message
	// *Note*: PartitionAny to let the broker handle load balancing across partitions
	message := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &s.Topic, Partition: kafka.PartitionAny},
		Key:            []byte(id),
		Value:          payload,
	}
	// produce message in async way
	// *Note*: Delivery reports are handled by the background event loop
	err = s.Producer.Produce(message, nil)
	// did we get an error?
	if err != nil {
		return fmt.Errorf("failed to produce message: %w", err)
	}
	return nil
}

// Close flushes any pending messages to the Kafka broker and shuts
// down the producer. It allows a grace period of 15 seconds for
// outstanding messages to be delivered before terminating connection.
func (s *KafkaStore) Close() {
	// wait 15 seconds for outstanding messages to send
	s.Producer.Flush(15 * 1000)
	// release underlying network resources and close the producer
	s.Producer.Close()
}

// executeEventLoop processes delviery reports and internal producer events
// It is intended to run as a background goroutine for the duration of the
// store's lifecycle.
func (s *KafkaStore) executeEventLoop() {
	// monitor the producer's events channel for delivery feedback and errors.
	// it is closed when the producer is shut down via the Close() method.
	for e := range s.Producer.Events() {
		switch ev := e.(type) {
		case *kafka.Message:
			// process message delivery reports
			// confirms whether kafka service or the broker acknowledged the message
			if ev.TopicPartition.Error != nil {
				fmt.Printf("Delivery failed:\t %v \n", ev.TopicPartition.Error)
			} else {
				fmt.Printf("Delivered to:\t %v \n", ev.TopicPartition)
			}
		case *kafka.Error:
			// handle internal kafka client errors.
			fmt.Printf("Producer Error:\t %v \n", ev)
		}
	}
}
