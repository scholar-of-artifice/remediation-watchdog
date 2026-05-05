package worker

import (
	"context"
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type Worker struct {
	Consumer  *kafka.Consumer // which service is consuming events?
	Processor *Processor      // what is handling these events?
	Topic     string          // what is the topic being handled?
}

// NewWorker initializes the Kafka consumer with a specific GroupID
// The GroupID is essential for Kafka to track what messages have already been processed.
func NewWorker(addr, groupID, topic string, proc *Processor) (*Worker, error) {
	//
	kafkaConfig := &kafka.ConfigMap{
		"bootstrap.servers": addr,
		"group.id":          groupID,
		"auto.offset.reset": "earliest", // start from beginning if no offset exists
	}
	// try to make a consumer
	c, err := kafka.NewConsumer(kafkaConfig)
	// did we get an error?
	if err != nil {
		return nil, fmt.Errorf("Failed to create consumer: %w", err)
	}
	//
	return &Worker{
		Consumer:  c,
		Processor: proc,
		Topic:     topic,
	}, nil
}

func (w *Worker) Start(ctx context.Context) {
	// try to subscribe to the topic
	err := w.Consumer.SubscribeTopics([]string{w.Topic}, nil)
	// did we get an error?
	if err != nil {
		log.Fatalf("Failed to subscribe to topic: %v", err)
	}
	log.Printf("eager-marmot worker started. Consuming events from %s...", w.Topic)
	for {
		select {
		case <-ctx.Done():
			// base case... we exit
			log.Println("Worker shutting down: Context cancelled")
			return
		default:
			// poll kafka for new messages
			w.pollForEvents(ctx)
		}
	}
}

// Close ensures the consumer leaves the group and closes network connections
func (w *Worker) Close() {
	log.Println("Closing Kafka consumer...")
	w.Consumer.Close()
}

func (w *Worker) pollForEvents(ctx context.Context) {
	// poll kafka for new messages with a 100ms timeout
	ev := w.Consumer.Poll(100)
	// did we get an event?
	if ev == nil {
		return
	}
	switch e := ev.(type) {
	case *kafka.Message:
		w.Processor.ProcessMessage(ctx, e)
	case kafka.Error:
		log.Printf("Kafka Error: %v", e)
	}
}
