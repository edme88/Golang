package main

import (
	"apirest/handlers" //actualizar
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main(){
	//Rutas
	mux := mux.NewRouter()

	//Endpoint
	mux.HandleFunc("/api/users/", handlers.GetUsers).Methods("GET")
	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.GetUser).Methods("GET")
	mux.HandleFunc("/api/user/", handlers.CreateUser).Methods("POST")
	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.UpdateUser).Methods("PUT")
	mux.HandleFunc("/api/user/{id:[0-9]+}", handlers.DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000",mux))
}