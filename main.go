package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

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

	// Views voor pagina's
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/overview", overviewHandler)
	r.HandleFunc("/portal", portalHandler)
	r.HandleFunc("/settle", settleHandler)

	// Views voor acties
	r.HandleFunc("/register", registerHandler).Methods("POST")

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

// Handler voor de portal pagina
func portalHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "portal.html", nil)
}

// Handler voor de settle pagina
func settleHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "settle.html", nil)
}

// Handler voor de overview pagina
func registerHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	mail := r.PostForm.Get("mail")
	name := r.PostForm.Get("name")
	password := r.PostForm.Get("password")
	passwordRepeat := r.PostForm.Get("passwordRepeat")
	fmt.Println(mail, name, password, passwordRepeat)
	http.Redirect(w, r, "/overview", 302)
}
