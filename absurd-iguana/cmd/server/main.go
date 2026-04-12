package main

import (
	"absurd-iguana/internal/api"
	"absurd-iguana/internal/store"
	"log"
	"net/http"
)

func main() {
	//initialize redis
	redisAddr := "localhost:6379" // use env vars for kubernets later
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
