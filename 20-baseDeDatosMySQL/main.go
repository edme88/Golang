package main

import (
	"fmt"
	"mysql/db"     //Lo cambiamos del original por el nombre que tuvo - Warning porque detecta varias carpetas. Se puede solucionar abriendo solo esta carpeta de 20-baseDeDatosMySQL
	"mysql/models" //lo cambiamos del original
)

func main(){
	db.Connect()

	fmt.Println(db.ExistsTable("users"))
	//db.CreateTable(models.UserSchema, "users")
	models.CreateUser("agusNew", "agusN123", "agusNew@ejemplo.com")
	//db.Ping()
	//db.TruncateTable("users")
	//user := models.GetUser(2)
	//fmt.Println(user)

	// user.Username = "Juan"
	// user.Password = "juan789"
	// user.Email = "juan@email.com"
	// user.Save()

	//user.Delete()

	//db.TruncateTable("users")
	users := models.ListUsers()
	
	fmt.Println(users)
	db.Close()
}