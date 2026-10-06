// Command greeter runs the greeter HTTP service.
package main

import (
	"log"
	"net/http"
	"os"

	"greeter/internal/handlers"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /hello", handlers.Hello)

	addr := ":" + port
	log.Printf("greeter listening on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("greeter: %v", err)
	}
}
