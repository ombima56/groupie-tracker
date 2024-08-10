package main

import (
	"log"
	"net/http"

	"groupie-trcker/handlers"
)

func main() {
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/artists", handlers.GetArtistsHandler)
	http.HandleFunc("/locations", handlers.GetLocationsHandler)
	http.HandleFunc("/dates", handlers.GetDatesHandler)
	http.HandleFunc("/relations", handlers.GetRelationsHandler)
	http.HandleFunc("/artist", handlers.GetArtistByIDHandler)

	log.Println("Server starting on port http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Printf("Could not start server: %s\n", err)
	}
}
