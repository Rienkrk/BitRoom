package main

import (
	"fmt"
	"html/template"
	"net/http"
	"github.com/gorilla/mux"
)

// Maak de templates voor alle functies bereikbaar
var templates *template.Template

func main() {

	// Parse de templates
	templates = template.Must(template.ParseGlob("templates/*.html"))

	// Initialiseer de router
	r := mux.NewRouter()

	// Vertel de server waar de static files zijn
	staticFileDirectory := http.Dir("./static/")
	staticFileHandler := http.StripPrefix("/static/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/static/").Handler(staticFileHandler).Methods("GET")

	// Views voor pagina's
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/login", loginHandler)
	r.HandleFunc("/overview", overviewHandler)
	r.HandleFunc("/portal", portalHandler)
	r.HandleFunc("/settle", settleHandler)

	// Views voor acties
	r.HandleFunc("/register", registerHandler).Methods("POST")
	r.HandleFunc("/addEntry", entryHandler).Methods("POST")

	// Initialiseer de server
	http.Handle("/", r)
	http.ListenAndServe(":8080", r)

}

// Handler voor entries
func entryHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/overview", 302)
}

// Handler voor de landings pagina
func indexHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", nil)
}

// Handler voor de login pagina
func loginHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "login.html", nil)
}

// Handler voor de overzicht pagina
func overviewHandler(w http.ResponseWriter, r *http.Request) {
		templates.ExecuteTemplate(w, "overview.html", nil)
}

// Handler voor de portaal pagina
func portalHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "portal.html", nil)
}

// Handler voor de vereken pagina
func settleHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "settle.html", nil)
}

// Handler voor het registreren
func registerHandler(w http.ResponseWriter, r *http.Request) {

	// Parse de form
	r.ParseForm()

	// Haal data uit de form
	mail := r.PostForm.Get("mail")
	name := r.PostForm.Get("name")
	password := r.PostForm.Get("password")
	passwordRepeat := r.PostForm.Get("passwordRepeat")

	// Print naar console
	fmt.Println(mail, name, password, passwordRepeat)

	// Redirect naar de overview pagina
	http.Redirect(w, r, "/portal", 302)
}
