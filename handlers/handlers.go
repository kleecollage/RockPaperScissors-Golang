package handlers

import (
	"log"
	"net/http"
	"text/template"
)

type Player struct {
	Name string
}

var player Player

const (
	templateDir  = "templates/"
	templateBase = templateDir + "base.html"
)

func Index(w http.ResponseWriter, r *http.Request) {
	restartValue()
	renderTemplate(w, "index.html", nil)
}

func NewGame(w http.ResponseWriter, r *http.Request) {
	restartValue()
	renderTemplate(w, "new-game.html", nil)
}

func Game(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error al parsear formulario", http.StatusBadRequest)
			return
		}
		player.Name = r.Form.Get("name")
	}
	// fmt.Println(player.Name) // debug
	if player.Name == "" {
		http.Redirect(w, r, "/new", http.StatusFound)
	}

	renderTemplate(w, "game.html", player)
}

func Play(w http.ResponseWriter, r *http.Request) {
	// TODO template
	renderTemplate(w, "index.html", nil)
}

func About(w http.ResponseWriter, r *http.Request) {
	restartValue()
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

func restartValue() {
	player.Name = ""
}
