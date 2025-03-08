package main

import (
	"fmt"
	"net/http"
)

func main() {
	//recibe una url y una función handler
	//para registrar un controlador
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hola mundo!")
	})

	port := ":8080"
	fmt.Printf("Servidor escuchando en http://localhost%s\n", port)
	http.ListenAndServe(port, nil)

}