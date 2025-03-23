package handlers

import (
	"fmt"
	"net/http"
)

func GetUsers(rw http.ResponseWriter, r *http.Request){
	fmt.Println(rw, "Listar todos los usuarios")
}

func GetUser(rw http.ResponseWriter, r *http.Request){
	fmt.Println(rw, "Obtiene un susuario")
}

func CreateUser(rw http.ResponseWriter, r *http.Request){
	fmt.Println(rw, "Crea un susuario")
}

func UpdateUser(rw http.ResponseWriter, r *http.Request){
	fmt.Println(rw, "Actualiza un susuario")
}

func DeleteUser(rw http.ResponseWriter, r *http.Request){
	fmt.Println(rw, "Borra un susuario")
}