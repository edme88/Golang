package main

import (
	"log"
	"net/http"
	"rps/handlers" //rock-paper-scisor
)

func main() {
	//Router: Componente que se encarga de redirigir la solicitud http
	//a la funcón handler correspondiente
	//Crear enrutador
	router := http.NewServeMux()

	//Configurar rutas
	router.HandleFunc("/", handlers.Index)
	router.HandleFunc("/new", handlers.NewGame)
	router.HandleFunc("/game", handlers.Game)
	router.HandleFunc("/play", handlers.Play)
	router.HandleFunc("/about", handlers.About)

	port := ":8080"
	log.Printf("Servidor escuchando en http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))

}