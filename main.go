package main

import (
	"fmt"
	"log"
	"net/http"
	"stox/router"
)

func main() {
	// Create a router instance.
	r := router.Router()

	// Set up the HTTP server.
	port := ":8080"
	fmt.Printf("Starting Server on port %s...\n", port)

	// ListenAndServe returns an error.
	err := http.ListenAndServe(port, r)
	if err != nil {
		log.Printf("Failed to start the server: %v\n", err)
		fmt.Println("Exiting...")
		return
	}
}
