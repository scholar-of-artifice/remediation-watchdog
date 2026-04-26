package main

import (
	"absurd-iguana/internal/api"
	"absurd-iguana/internal/store"
	"log"
	"net/http"
	"os"
)

func main() {
	// look for environment variables
	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "localhost:6379" // maybe cause error instead?
	}
	//initialize redis
	storage := store.NewRedisStore(redisAddr)
	h := &api.Handler{Repo: storage}

	//route defintion
	mux := http.NewServeMux()
	mux.HandleFunc("POST /events", h.ProduceEventHandler)
	log.Println("absurd-iguana starting on :8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
