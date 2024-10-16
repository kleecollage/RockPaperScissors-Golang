package handlers

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Index is working!")
}

func NewGame(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Crear nuevo juego")
}

func Game(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Game is working!")
}

func Play(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Play is working!")
}

func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "About is working!")
}
