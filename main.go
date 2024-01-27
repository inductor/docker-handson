package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Handler function to respond to HTTP requests
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	// Start the server on port 8080 and log any errors
	fmt.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
