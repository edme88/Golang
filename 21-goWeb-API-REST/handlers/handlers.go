package handlers

import (
	"apirest/db"     //actualizar
	"apirest/models" //actualizar
	"encoding/json"
	"fmt"
	"net/http"
)

func GetUsers(rw http.ResponseWriter, r *http.Request){
	rw.Header().Set("Content-Type","application/json")
	db.Connect()
	users := models.ListUsers()
	db.Close()
	output, _ := json.Marshal(users)
	//fmt.Println(rw, string(output))
	//rw.Write(output)
	fmt.Fprintln(rw, string(output))
}

func GetUser(rw http.ResponseWriter, r *http.Request){
	fmt.Fprintln(rw, "Obtiene un susuario")
}

func CreateUser(rw http.ResponseWriter, r *http.Request){
	fmt.Fprintln(rw, "Crea un susuario")
}

func UpdateUser(rw http.ResponseWriter, r *http.Request){
	fmt.Fprintln(rw, "Actualiza un susuario")
}

func DeleteUser(rw http.ResponseWriter, r *http.Request){
	fmt.Fprintln(rw, "Borra un susuario")
}