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

type TemplateData struct {
	Artists   []models.Artist
	Query     string
	NoResults bool
}

type ArtistDetailData struct {
	Artist   models.Artist
	Relation models.Relation
}

func ErrorHandler(w http.ResponseWriter, message string, statusCode int) {
	// Set the status code
	w.WriteHeader(statusCode)
	// Define error template data
	data := struct {
		StatusCode int
		ErrMsg     string
	}{
		StatusCode: statusCode,
		ErrMsg:     message,
	}
	tmpl, err := template.ParseFiles("templates/error.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, data); err != nil {
		log.Println("Error executing error template")
		http.Error(w, "Error executing data deatils", http.StatusInternalServerError)
		return
	}
}

func ServeArtists(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	artists, err := services.GetArtists()
	if err != nil {
		log.Printf("Error getting artists: %v", err)
		ErrorHandler(w, "Unable to retrieve artists at this time. Please try again later.", http.StatusInternalServerError)
		return
	}
	filteredArtists := filterArtists(artists, query)

	// Only show the "no results" error if the query is not empty
	if len(filteredArtists) == 0 && query != "" {
		ErrorHandler(w, "We couldn't find any artists matching your search criteria. Please try a different term or check your spelling.", http.StatusNotFound)
		return
	}

	data := TemplateData{
		Artists:   filteredArtists,
		Query:     query,
		NoResults: len(filteredArtists) == 0 && query != "",
	}

	tmpl, err := template.ParseFiles("templates/artists.html")
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		ErrorHandler(w, "An unexpected error occurred. Please try again later.", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		ErrorHandler(w, "We encountered an issue while rendering the page. Please try again later.", http.StatusInternalServerError)
		return
	}
}

// filterArtists filters the list of artists based on the search query
func filterArtists(artists []models.Artist, query string) []models.Artist {
	if query == "" {
		return artists
	}
	var result []models.Artist
	for _, a := range artists {
		if strings.Contains(strings.ToLower(a.Name), strings.ToLower(query)) {
			result = append(result, a)
		}
	}
	return result
}

func ServeArtistDetails(w http.ResponseWriter, r *http.Request) {
	// Check if the path is just "/artist/", which is not a valid artist page
	if r.URL.Path == "/artist/" || len(r.URL.Path) <= len("/artist/") {
		ErrorHandler(w, "Artist not found", http.StatusNotFound)
		return
	}

	// Extract the artist ID from the URL path
	idStr := r.URL.Path[len("/artist/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		ErrorHandler(w, "Invalid artist ID", http.StatusBadRequest)
		return
	}

	// Get the artist and relation data
	artist, relation, err := services.GetArtistByID(id)
	if err != nil {
		if err.Error() == "artist not found" {
			ErrorHandler(w, "Artist not found", http.StatusNotFound)
		} else {
			log.Printf("Error retrieving artist by ID %v: %s", id, err)
			ErrorHandler(w, "Unable to retrieve artist details at this time. Please try again later.", http.StatusInternalServerError)
		}
		return
	}

	// Prepare the template data
	data := ArtistDetailData{
		Artist:   *artist,
		Relation: *relation,
	}

	// Render the artist details template
	tmpl, err := template.ParseFiles("templates/artist_details.html")
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		ErrorHandler(w, "Unable to load artist details at this time. Please try again later.", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		ErrorHandler(w, "Error rendering artist details. Please try again later.", http.StatusInternalServerError)
		return
	}
}

// GetArtistsHandler handles the /artists route
func GetArtistsHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	artists, err := services.GetArtists()
	if err != nil {
		log.Printf("Error fetching artists: %v", err)
		ErrorHandler(w, "Unable to retrieve artist information at this time. Please try again later.", http.StatusInternalServerError)
		return
	}
	filteredArtists := filterArtists(artists, query)
	if len(filteredArtists) == 0 {
		ErrorHandler(w, "No artists found matching the search term.", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(filteredArtists); err != nil {
		log.Printf("Error encoding artists data to JSON: %v", err)
		ErrorHandler(w, "An error occurred while processing the artist data. Please try again later.", http.StatusInternalServerError)
		return
	}
}

// GetLocationsHandler handles the /locations route
func GetLocationsHandler(w http.ResponseWriter, r *http.Request) {
	locations, err := services.GetLocations()
	if err != nil {
		log.Printf("Error fetching locations: %v", err)
		ErrorHandler(w, "Unable to retrieve locations at this time. Please try again later.", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(locations); err != nil {
		log.Printf("Error encoding locations data to JSON: %v", err)
		ErrorHandler(w, "An error occurred while processing location data. Please try again later.", http.StatusInternalServerError)
		return
	}
}

// GetDatesHandler handles the /dates route
func GetDatesHandler(w http.ResponseWriter, r *http.Request) {
	dates, err := services.GetDates()
	if err != nil {
		log.Printf("Error fetching dates: %v", err)
		ErrorHandler(w, "Unable to retrieve dates at this time. Please try again later.", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(dates); err != nil {
		log.Printf("Error encoding dates data to JSON: %v", err)
		ErrorHandler(w, "An error occurred while processing date information. Please try again later.", http.StatusInternalServerError)
		return
	}
}

// GetRelationsHandler handles the /relations route
func GetRelationsHandler(w http.ResponseWriter, r *http.Request) {
	relations, err := services.GetRelations()
	if err != nil {
		log.Printf("Error fetching relations: %v", err)
		ErrorHandler(w, "Unable to retrieve relations at this time. Please try again later.", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(relations); err != nil {
		log.Printf("Error encoding relations data to JSON: %v", err)
		ErrorHandler(w, "An error occurred while processing relation data. Please try again later.", http.StatusInternalServerError)
		return
	}
}

func GetArtistByIDHandler(w http.ResponseWriter, r *http.Request) {
	// Get artist ID from the URL path
	idStr := r.URL.Path[len("/artist/"):]
	artistID, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("Invalid artist ID: %v", err)
		ErrorHandler(w, "Invalid artist ID provided. Please check and try again.", http.StatusBadRequest)
		return
	}
	artist, relation, err := services.GetArtistByID(artistID)
	if err != nil {
		// Check if the artist was not found
		if err.Error() == "artist not found" {
			ErrorHandler(w, "Artist not found. Please check the ID and try again.", http.StatusNotFound)
		} else {
			log.Printf("Error fetching artist or relation with ID %d: %v", artistID, err)
			ErrorHandler(w, "Unable to retrieve artist details at this time. Please try again later.", http.StatusInternalServerError)
		}
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
		log.Printf("Error encoding response for artist ID %d: %v", artistID, err)
		ErrorHandler(w, "An error occurred while processing the response. Please try again later.", http.StatusInternalServerError)
		return
	}
}
