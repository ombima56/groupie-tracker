package main

import (
	"log"
	"net/http"

	"groupie-trcker/routes"
)

func main() {
	routes.RegisterRoutes()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Server starting on port http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Printf("Could not start server: %s\n", err)
	}
}
