package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

//recibe una url y una función handler
//para registrar un controlador
func Index(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(w, "<h1>Página de Inicio!</h1>") //No usar html
	//tpl, err := template.ParseFiles("templates/index.html") //Analiza una lista del nombre de las plantillas
	tpl, err := template.ParseFiles("templates/base.html","templates/index.html")
	if err != nil {
		http.Error(w, "Error al analizar plantillas", http.StatusInternalServerError)
	}

	// //Para enviar datos se pueden emplear estructuras o mapas
	// //de manera directa crear estrictura e inicializar los datos
	// data := struct {
	// 	Title string
	// 	Message string
	// }{
	// 	Title: "Página de inicio",
	// 	Message: "Bienvenido a Pidra, Papel o Tijera",
	// }
	// err = tpl.Execute(w, data)
	// if err != nil{
	// 	http.Error(w, "Error al renderizar la plantilla", http.StatusInternalServerError)
	// }
	err = tpl.ExecuteTemplate(w, "base", nil)
	if err != nil{
		http.Error(w, "Error al renderizar la plantilla", http.StatusInternalServerError)
	}
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