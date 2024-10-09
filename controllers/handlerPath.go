package controllers

import (
	"net/http"
)

func HandlerPath(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		ServeArtists(w, r)

	case "/static/":
		ErrorHandler(w, "Access forbidden. Please check your URL or navigate to the homepage.", http.StatusForbidden, false, false)
	default:
		ErrorHandler(w, "Page Not Found", http.StatusNotFound, true, true)
	}
}
