package main

import (
	"chat-websockets/internal/handlers"
	"log"
	"net/http"
)

func main() {
	mux := routes()
	log.Println("Starting web server on port 8080")
	go handlers.ListenToWsChannel()

	_ = http.ListenAndServe(":8080", mux)
}
