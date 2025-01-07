package app

import (
	"log"
	"net/http"
)

func StartServer() {

	// Define the port to listen on
	port := ":8080"

	// Start the HTTP server
	log.Printf("Starting HTTP server on port %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}

}
