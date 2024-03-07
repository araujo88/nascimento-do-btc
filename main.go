package main

import (
	"log"
	"net/http"
	"strings"
)

func main() {
	// Define a new file server handler that wraps http.FileServer
	fileServer := http.FileServer(http.Dir("./static"))

	// Create a custom handler to wrap around the file server
	loggedFileServer := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log every request
		log.Printf("Received request for %s from %s", r.URL.Path, r.RemoteAddr)

		// Prevent directory traversal attack by cleaning the path and checking for attempts to escape the root
		if strings.Contains(r.URL.Path, "..") {
			log.Printf("Attempted directory traversal: %s", r.URL.Path)
			http.Error(w, "Access denied!", http.StatusForbidden)
			return
		}

		// Serve the file
		fileServer.ServeHTTP(w, r)
	})

	// Use the wrapped handler to serve static files
	http.Handle("/", loggedFileServer)

	// Start the server on port 80 and log any errors
	log.Println("Listening on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
