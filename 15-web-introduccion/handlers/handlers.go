package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"rps/rps" //package
	"strconv"
)

type Player struct {
	Name string
}

var player Player //instancia del jugador

const (
	templateDir = "templates/"
	templateBase = templateDir + "base.html"
)

//recibe una url y una función handler
//para registrar un controlador
//func Index(w http.ResponseWriter, r *http.Request) {
//fmt.Fprintln(w, "<h1>Página de Inicio!</h1>") //No usar html
//tpl, err := template.ParseFiles("templates/index.html") //Analiza una lista del nombre de las plantillas
// tpl, err := template.ParseFiles("templates/base.html","templates/index.html")
// if err != nil {
// 	http.Error(w, "Error al analizar plantillas", http.StatusInternalServerError)
// }

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
// 	err = tpl.ExecuteTemplate(w, "base", nil)
// 	if err != nil{
// 		http.Error(w, "Error al renderizar la plantilla", http.StatusInternalServerError)
// 	}
// }

func Index(w http.ResponseWriter, r *http.Request) {
	restartValue()
	renderTemplate(w, "index.html", nil)
}

func NewGame(w http.ResponseWriter, r *http.Request) {
	restartValue()
	renderTemplate(w, "new-game.html", nil)
}

func Game(w http.ResponseWriter, r *http.Request) {
	//procesador los datos del formulario
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form", http.StatusBadRequest)
			return
		}

		player.Name = r.Form.Get("name")
	}

	if player.Name == "" {
		http.Redirect(w, r, "/new", http.StatusFound)
	}

	fmt.Println(player.Name)
	renderTemplate(w, "game.html", player)
}

func Play(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(w, "Jugar")
	playerChoice, _ := strconv.Atoi(r.URL.Query().Get("c"))
	result := rps.PlayRound(playerChoice)

	//Ingresar a http://localhost:8080/play?c=1
	//fmt.Println(result)

	out, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)

	//Ingresar a http://localhost:8080/play?c=2 y ver que pasa

}

func About(w http.ResponseWriter, r *http.Request) {
	restartValue()
	renderTemplate(w, "about.html", nil)
}

func renderTemplate(w http.ResponseWriter, page string, data any){
	tpl := template.Must(template.ParseFiles(templateBase, templateDir+page))

	err := tpl.ExecuteTemplate(w, "base", data)
	if err != nil{
		http.Error(w, "Error al renderizar la plantilla", http.StatusInternalServerError)
		log.Println(err)
		return
	}

}

//Reiniciar valores
func restartValue(){
	player.Name = ""
	rps.ComputerScore = 0
	rps.PlayerScore = 0
} 