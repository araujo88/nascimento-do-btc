package main

import (
	"log"
	"net/http"
)

func main() {
	// Set the directory containing your static files.
	fs := http.FileServer(http.Dir("./static"))

	// Serve static files.
	http.Handle("/", fs)

	// Start the server on port 80.
	log.Println("Listening on :80...")
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}
