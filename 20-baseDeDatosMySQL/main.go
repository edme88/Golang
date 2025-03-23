package main

import (
	"fmt"
	"mysql/db"     //Lo cambiamos del original por el nombre que tuvo - Warning porque detecta varias carpetas. Se puede solucionar abriendo solo esta carpeta de 20-baseDeDatosMySQL
	"mysql/models" //lo cambiamos del original
)

func main(){
	db.Connect()

	fmt.Println(db.ExistsTable("users"))
	db.CreateTable(models.UserSchema, "users")
	db.Ping()
	db.Close()
}