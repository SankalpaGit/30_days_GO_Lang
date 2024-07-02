package main

import (
	"html/template"
	"net/http"
	"sync"
)

var (
	tmplCache = template.Must(template.ParseGlob("templates/*.html"))
	books     = []string{}
	mu        sync.Mutex
)

func main() {
	// Serve static files from the "static" directory
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Define handlers for dynamic content
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/books", bookListHandler)
	http.HandleFunc("/addbook", addBookHandler)

	// Start the server
	http.ListenAndServe(":8080", nil)
}

// Handlers
func homeHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.html", nil)
}

func bookListHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	renderTemplate(w, "booklist.html", books)
}

func addBookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		title := r.FormValue("title")
		if title != "" {
			mu.Lock()
			books = append(books, title)
			mu.Unlock()
		}
		http.Redirect(w, r, "/books", http.StatusSeeOther)
		return
	}
	renderTemplate(w, "addbook.html", nil)
}

// Helper function to render templates
func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	err := tmplCache.ExecuteTemplate(w, tmpl, data)
	if err != nil {
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
	}
}
