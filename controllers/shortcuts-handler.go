package controllers

import (
	"log"
	"net/http"
	"text/template"
)

func ServeShortcuts(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/shortcuts" {
		if r.Method != http.MethodGet {
			ErrorHandler(w, "method not allowed", http.StatusMethodNotAllowed, true, true)
			return
		}
	}

	tmpl, err := template.ParseFiles("templates/shortcuts.html")
	if err != nil {
		ErrorHandler(w, "file not found", http.StatusNotFound, true, true)
		log.Printf("Error parsing index.html: %v\n", err)
		return
	}
	
	err = tmpl.Execute(w, nil)
	if err != nil {
		ErrorHandler(w, "Something unexpected occured", http.StatusInternalServerError, false, false)
		log.Printf("Error executing template: %v\n", err)
		return
	}
}
