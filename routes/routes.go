package routes

import (
	"net/http"

	"groupie-trcker/handlers"
)

func RegisterRoutes() {
	http.HandleFunc("/artists", handlers.GetArtistsHandler)
	http.HandleFunc("/artist/", handlers.ServeArtistDetails)
	http.HandleFunc("/locations", handlers.GetLocationsHandler)
	http.HandleFunc("/dates", handlers.GetDatesHandler)
	http.HandleFunc("/relations", handlers.GetRelationsHandler)
}
