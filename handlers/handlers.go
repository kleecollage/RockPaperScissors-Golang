package handlers

import (
	"log"
	"net/http"
	"text/template"
)

const (
	templateDir  = "templates/"
	templateBase = templateDir + "base.html"
)

func Index(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html", nil)
}

func NewGame(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "new-game.html", nil)
}

func Game(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "game.html", nil)
}

func Play(w http.ResponseWriter, r *http.Request) {
	// TODO template
	renderTemplate(w, "index.html", nil)
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.html", nil)
}

func renderTemplate(w http.ResponseWriter, page string, data any) {
	tmpl := template.Must(template.ParseFiles(templateBase, templateDir+page))
	err := tmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, "Error al renderizar plantillas", http.StatusInternalServerError)
		log.Println(err)
		return
	}
}
