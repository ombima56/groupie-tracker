package routes

import (
	"net/http"

	"groupie-trcker/handlers"
)

func RegisterRoutes() {
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/artists", handlers.GetArtistsHandler)
	http.HandleFunc("/locations", handlers.GetLocationsHandler)
	http.HandleFunc("/dates", handlers.GetDatesHandler)
	http.HandleFunc("/relations", handlers.GetRelationsHandler)
	http.HandleFunc("/artist/", handlers.ServeArtistDetails)
}
