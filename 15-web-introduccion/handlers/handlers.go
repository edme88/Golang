package handlers

import (
	"fmt"
	"net/http"
)

//recibe una url y una función handler
//para registrar un controlador
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Página de Inicio!")
}

func NewGame(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Crear nuevo juego")
}

func Game(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Juego")
}

func Play(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Jugar")
}

func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Acerca de")
}