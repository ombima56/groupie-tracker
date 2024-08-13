package main

import (
	"groupie-trcker/handlers"
	"groupie-trcker/routes"

	"log"
	"net/http"
	"strings"
)

func main() {
	routes.RegisterRoutes()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// Catch-all for undefined routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Check if the request is for the root path or favicon.ico
		if r.URL.Path == "/" || r.URL.Path == "/favicon.ico" {
			handlers.ServeArtists(w, r)
			return
		}
		// Check if the request is for an artist details page
		if strings.HasPrefix(r.URL.Path, "/artist/") {
			handlers.ServeArtistDetails(w, r)
			return
		}
		// If it's not a known route, use ErrorHandler for 404
		handlers.ErrorHandler(w, "404 Page Not Found", http.StatusNotFound)
	})

	log.Println("Server starting on port http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Printf("Could not start server: %s\n", err)
	}
}
