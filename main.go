package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"learn.zone01kisumu.ke/git/johnodhiambo0/groupie-tracker-visualizations/controllers"
)

func main() {
	if len(os.Args) != 1 {
		fmt.Println("Incorrect number of arguments passed. Usage: go run .")
		return
	}

	http.Handle("/static/", controllers.CustomFileServer("./static"))
	http.HandleFunc("/artists", controllers.ServeArtists)
	http.HandleFunc("/artist/", controllers.ServeArtistDetails)
	http.HandleFunc("/about", controllers.AboutHandler)
	http.HandleFunc("/search-suggestions", controllers.GetSearchSuggestionsHandler)
	http.HandleFunc("/", controllers.HandlerPath)

	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
