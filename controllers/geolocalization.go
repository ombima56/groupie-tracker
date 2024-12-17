package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

var coordinates map[string][]float64

func sanitizeLocation(location string) string {
	// Replace underscores and hyphens with spaces
	location = strings.ReplaceAll(location, "_", " ")
	location = strings.ReplaceAll(location, "-", " ")
	return location
}

// get coordinates from location using mapbox api
func Geocode(location string) (float64, float64, error) {
	apiKey := "pk.eyJ1IjoiYnJhdm4iLCJhIjoiY20ydmJpN3dvMGRjdTJpcXl4bGd6bHJpeiJ9.mlEfzqY52K9hUCr9KvyIGQ"
	location = sanitizeLocation(location)
	url := fmt.Sprintf("https://api.mapbox.com/geocoding/v5/mapbox.places/%s.json?access_token=%s", location, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return 0, 0, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, 0, err
	}
	var data struct {
		Features []struct {
			Geometry struct {
				Coordinates []float64 `json:"coordinates"`
			} `json:"geometry"`
		} `json:"features"`
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return 0, 0, err
	}
	if len(data.Features) == 0 {
		return 0, 0, fmt.Errorf("no coordinates found for location: %s", location)
	}
	coordinates := data.Features[0].Geometry.Coordinates
	return coordinates[1], coordinates[0], nil
}

// function to send coordinates to map.js as json
func SendCoordinates(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("coordinates:", coordinates)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(coordinates)
}
