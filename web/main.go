package main

import (
	"html/template"
	"net/http"
	"path/filepath"
)

// HomeHandler serves the home.html file
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	htmlPath := filepath.Join("templates", "home.html")
	tmpl, err := template.ParseFiles(htmlPath)
	if err != nil {
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Unable to render template", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/home", HomeHandler)
	http.ListenAndServe(":8080", nil)
}
