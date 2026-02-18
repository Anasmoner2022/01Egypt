package main

import (
	"fmt"
	"groupi_tracker/handlers"
	"log"
	"net/http"
)

func main() {
	// Serve static files (CSS, JS)
	// FIX: Must be "/static/" with trailing slash for FileServer to work
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Register handlers
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/artist", handlers.ArtistHandler)

	// NEW: Search/filter feature (client-server interaction requirement)
	http.HandleFunc("/search", handlers.SearchHandler)

	// Start server
	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
