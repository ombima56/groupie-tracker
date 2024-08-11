package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"text/template"

	"groupie-trcker/models"
	"groupie-trcker/services"
)

type templateData struct {
	Artist []models.Artist
}

type ArtistDetailData struct {
	Artist   models.Artist
	Relation models.Relation
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "page not found", http.StatusNotFound)
		log.Println("wewe utakufa vibaya wewe")
		return
	}

	artist, err := services.GetArtists()
	if err != nil {
		http.Error(w, "Unexpected error occures, try again later", http.StatusInternalServerError)
		log.Printf("Failed to get artists: %v", err)
		return
	}

	Data := templateData{Artist: artist}

	tmpl, err := template.ParseFiles("templates/artists.html")
	if err != nil {
		http.Error(w, "Unexpected error occures, try again later", http.StatusInternalServerError)
		log.Printf("failed to parse index template: %v", err)
		return
	}

	err = tmpl.Execute(w, Data)
	if err != nil {
		http.Error(w, "Unexpected error occures, try again later", http.StatusInternalServerError)
		log.Printf("failed to execute template: %v", err)
		return
	}
}

// GetArtistsHandler handles the /artists route
func GetArtistsHandler(w http.ResponseWriter, r *http.Request) {
	artists, err := services.GetArtists()
	if err != nil {
		http.Error(w, "Failed to fetch artists", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(artists); err != nil {
		http.Error(w, "Failed to encode locations", http.StatusInternalServerError)
	}
}

// GetLocationsHandler handles the /locations route
func GetLocationsHandler(w http.ResponseWriter, r *http.Request) {
	locations, err := services.GetLocations()
	if err != nil {
		http.Error(w, "Failed to fetch locations", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(locations); err != nil {
		http.Error(w, "Failed to encode locations", http.StatusInternalServerError)
	}
}

// GetDatesHandler handles the /dates route
func GetDatesHandler(w http.ResponseWriter, r *http.Request) {
	dates, err := services.GetDates()
	if err != nil {
		http.Error(w, "Failed to fetch dates", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(dates); err != nil {
		http.Error(w, "Failed to encode locations", http.StatusInternalServerError)
	}
}

// GetRelationsHandler handles the /relations route
func GetRelationsHandler(w http.ResponseWriter, r *http.Request) {
	relations, err := services.GetRelations()
	if err != nil {
		http.Error(w, "Failed to fetch relations", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(relations); err != nil {
		http.Error(w, "Failed to encode locations", http.StatusInternalServerError)
	}
}

func GetArtistByIDHandler(w http.ResponseWriter, r *http.Request) {
	trimmedPath := strings.TrimSuffix(r.URL.Path, "/")

	// Get artist ID from the URL path
	idStr := strings.TrimPrefix(trimmedPath, "/artists/")

	artistID, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("Error converting ID:", err)
		http.Error(w, "Invalid artist ID", http.StatusBadRequest)
		return
	}

	artist, relation, err := services.GetArtistByID(artistID)
	if err != nil {
		log.Println("Error fetching artist or relation:", err)
		http.Error(w, "Failed to fetch artist or relation", http.StatusInternalServerError)
		return
	}

	// Create a response combining artist and relation data
	response := struct {
		Artist   *models.Artist   `json:"artist"`
		Relation *models.Relation `json:"relation"`
	}{
		Artist:   artist,
		Relation: relation,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Println("Error encoding response:", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func ServeArtistDetails(w http.ResponseWriter, r *http.Request) {
	if !strings.HasPrefix(r.URL.Path, "/artist/") {
		http.Error(w, "page not found", http.StatusNotFound)
		log.Printf("Invalid URL path")
		return
	}

	// Capture the artist ID from the URL path
	idStr := strings.TrimPrefix(r.URL.Path, "/artist/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid artist ID", http.StatusBadRequest)
		return
	}

	// Fetch the artist details
	artist, relation, err := services.GetArtistByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Prepare the data for the template
	data := ArtistDetailData{
		Artist:   *artist,
		Relation: *relation,
	}

	// Parse and execute the template
	tmpl, err := template.ParseFiles("templates/artist_details.html")
	if err != nil {
		http.Error(w, "Failed to parse artist details template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Failed to execute artist details template", http.StatusInternalServerError)
		log.Printf("Failed to execute artist details template: %v", err)
		return
	}
}
