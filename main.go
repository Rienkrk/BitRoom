package main

import "net/http"
import "github.com/gorilla/mux"
import "html/template"

// Maak de templates voor alle handles bereikbaar
var templates *template.Template

func main() {

	// Parse de templates
	templates = template.Must(template.ParseGlob("templates/*.html"))

	// Initialiseer de webserver
	r := mux.NewRouter()

	// Vertel de server waar de static files zijn
	staticFileDirectory := http.Dir("./static/")
	staticFileHandler := http.StripPrefix("/static/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/static/").Handler(staticFileHandler).Methods("GET")

	// Routes
	// r.HandleFunc("/overview", overviewHandler)
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/overview", overviewHandler)

	// Start de server
	http.Handle("/", r)
	http.ListenAndServe(":8080", r)

}

// Handler voor de index pagina
func indexHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", nil)
}

// Handler voor de overview pagina
func overviewHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "overview.html", nil)
}
