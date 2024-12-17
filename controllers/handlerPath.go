package controllers

import (
	"net/http"
	"path/filepath"
)

func HandlerPath(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		ServeArtists(w, r)

	case "/static/":
		ErrorHandler(w, "Access forbidden. Please check your URL or navigate to the homepage.", http.StatusForbidden, false, false)
	case "/coordinates":
		SendCoordinates(w, r)
	default:
		ErrorHandler(w, "Page Not Found", http.StatusNotFound, true, true)
	}
}

func CustomFileServer(dir string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the request is for the base directory
		if r.URL.Path == "/static/" {
			// Call the HandlerPath to handle the error
			HandlerPath(w, r)
			return
		}

		// Remove the "/static/" prefix to get the actual file path
		filePath := filepath.Join(dir, r.URL.Path[len("/static/"):])

		http.ServeFile(w, r, filePath)
	})
}
