package main

import (
	route "github.com/gocroot/route"
	"log"
	"net/http"
	"os"
)

func main() {
	// Handle the WebHook route
	http.HandleFunc("/", route.URL)

	// Get the server port from an environment variable (use SERVER_PORT, not MONGOSTRING)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on port: %s", port)

	// Start the server on the specified port
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
