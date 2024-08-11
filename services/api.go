package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"groupie-trcker/models"
)

// FetchData is a helper function to fetch data from the given URL and unmarshal it into the provided struct
func FetchData(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to fetch data from %s: %s", url, resp.Status)
	}

	return json.NewDecoder(resp.Body).Decode(target)
}

// DataFetcher interface to abstract the data fetching logic
type DataFetcher interface {
	FetchData(url string) ([]byte, error)
}

// Default Fetcher implementation with set timeout
var Fetcher DataFetcher = fetcher{
	client: &http.Client{
		Timeout: 20 * time.Second,
	},
}

// fetcher struct that implements the DataFetcher interface
type fetcher struct {
	client *http.Client
}

// FetchData makes an HTTP GET request to the given URL and returns the response body
func (f fetcher) FetchData(url string) ([]byte, error) {
	resp, err := f.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	return body, nil
}

// GetArtists fetches the artist data from the API and returns a slice of Artist structs
func GetArtists() ([]models.Artist, error) {
	body, err := Fetcher.FetchData(models.ArtistsURL)
	if err != nil {
		return nil, err
	}

	var artists []models.Artist
	if err := json.Unmarshal(body, &artists); err != nil {
		return nil, fmt.Errorf("failed to unmarshal artists: %v", err)
	}

	return artists, nil
}

// GetLocations fetches the location data from the API and returns a slice of Location structs
func GetLocations() ([]models.Location, error) {
	body, err := Fetcher.FetchData(models.LocationsURL)
	if err != nil {
		return nil, err
	}

	var locations struct {
		Index []models.Location `json:"index"`
	}

	if err := json.Unmarshal(body, &locations); err != nil {
		return nil, fmt.Errorf("failed to unmarshal locations: %v", err)
	}

	return locations.Index, nil
}

// GetDates fetches the date data from the API and returns a slice of Date structs
func GetDates() ([]models.Dates, error) {
	body, err := Fetcher.FetchData(models.DatesURL)
	if err != nil {
		return nil, err
	}

	var dates struct {
		Index []models.Dates `json:"index"`
	}

	if err := json.Unmarshal(body, &dates); err != nil {
		return nil, fmt.Errorf("failed to unmarshal dates: %v", err)
	}

	return dates.Index, nil
}

// GetRelations fetches the relation data from the API and returns a slice of Relation structs
func GetRelations() ([]models.Relation, error) {
	body, err := Fetcher.FetchData(models.RelationURL)
	if err != nil {
		return nil, err
	}

	var relations struct {
		Index []models.Relation `json:"index"`
	}
	if err := json.Unmarshal(body, &relations); err != nil {
		return nil, fmt.Errorf("failed to unmarshal relations: %v", err)
	}

	return relations.Index, nil
}

// GetArtistByID fetches the artist data by ID and returns the Artist struct along with its relation
func GetArtistByID(artistID int) (*models.Artist, *models.Relation, error) {
	// Fetch artist data
	artists, err := GetArtists()
	if err != nil {
		return nil, nil, err
	}

	// Find the artist with the specified ID
	var artist *models.Artist
	for _, a := range artists {
		if a.ID == artistID {
			artist = &a
			break
		}
	}
	if artist == nil {
		return nil, nil, fmt.Errorf("artist not found")
	}

	// Fetch relation data
	relations, err := GetRelations()
	if err != nil {
		return nil, nil, err
	}

	// Find the relation for the specific artist
	var relation *models.Relation
	for _, r := range relations {
		if r.ID == artistID {
			relation = &r
			break
		}
	}
	if relation == nil {
		return nil, nil, fmt.Errorf("relation not found for artist")
	}
	return artist, relation, nil
}
