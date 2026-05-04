package main

import (
	"absurd-iguana/internal/api"
	"absurd-iguana/internal/store"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// configuration lifecycle:
	// retrieve the kafka cluster address from environment variables.
	// failing here prevents the service from entering hanging state.
	kafkaAddr := os.Getenv("KAFKA_BOOTSTRAP_SERVERS")
	if kafkaAddr == "" {
		log.Fatal("KAFKA_BOOTSTRAP_SERVERS is required.")
	}
	// name the kafka topic
	topic := "remediation-events"
	// initialize the kafka producer
	kafkaStore, err := store.NewKafkaStore(kafkaAddr, topic)
	// did we get an error?
	if err != nil {
		log.Fatalf("Critical Failure: Could not connect to Kafka at %s: %v", kafkaAddr, err)
	}
	// ensure that all buffered messages are flushed to broker before the application terminates
	defer kafkaStore.Close()
	// inject the kafkaStore into the API handler
	h := &api.Handler{
		Repo: kafkaStore,
	}
	//route defintion
	mux := http.NewServeMux()
	mux.HandleFunc("/produce", h.ProduceEventHandler)
	// structred serveer start
	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	// create a channel to listen for OS signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// use named function in a goroutine
	go startServer(srv)
	// block main execution until the stop channel receives a signal
	<-stop // wait here until signal is received
	log.Println("Shutting down gracefully...")
	// give the server time to finsih
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Printf("Server forced to shutdown: %v", err)
	}
}

// startServer handles the execution of the HTTP server.
// It is non-exported to keep it scoped to this package.
func startServer(srv *http.Server) {
	log.Println("absurd-iguana starting on :8080...")
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Listen error: %v", err)
	}
}
