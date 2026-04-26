package store

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type KafkaStore struct {
	Producer *kafka.Producer
	Topic    string
}

// initializes the producer using the bootstrap server address
func NewKafkaStore(addr string, topic string) (*KafkaStore, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": addr,
		"client.id":         "absurd-iguana-producer",
		"acks":              "all",
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create producer: %w", err)
	}
	return &KafkaStore{
		Producer: p,
		Topic:    topic,
	}, nil
}

func (s *KafkaStore) SaveEvent(ctx context.Context, id string, data interface{}) error {
	// right now... just write directly to verify connection
	// convert data to json
	payload, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal event: %w", err)
	}
	// create the kafka message
	message := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &s.Topic, Partition: kafka.PartitionAny},
		Key:            []byte(id),
		Value:          payload,
	}
	// send the data to dazzling-remora
	err = &s.Producer.Produce(message, nil)
	if err != nil {
		return fmt.Errorf("failed to produce message: %w", err)
	}
	return nil
}

func (s *KafkaStore) Close() {
	// wait 15 seconds for outstanding messages to send
	s.Producer.Flush(time.Duration(15) * time.Second)
	// close the producer
	s.Producer.Close()
}
