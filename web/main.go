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
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/books", BookListHandler)
	http.HandleFunc("/addbook", AddBookHandler)

	http.ListenAndServe(":8080", nil)
}

// HomeHandler serves the home.html file
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.html", nil)
}

// BookListHandler serves the booklist.html file with the list of books
func BookListHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	renderTemplate(w, "booklist.html", books)
}

// AddBookHandler serves addbook.html for GET and processes form for POST
func AddBookHandler(w http.ResponseWriter, r *http.Request) {
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

// renderTemplate renders the specified template with the given data
func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	err := tmplCache.ExecuteTemplate(w, tmpl, data)
	if err != nil {
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
	}
}
