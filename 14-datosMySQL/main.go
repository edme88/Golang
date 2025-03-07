package main

import (
	"datosMySQL/database"
	"datosMySQL/handlers"

	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// dns := "root:GoLang123#@tcp(localhost:3306)/db_contacts" //datos de la BD para poder conectarnos usuario:Password@tcp(localhost:port)
	// //Abrir la conexión con la BD
	// db, err := sql.Open("mysql", dns) //Se utiliza el pacquete de forma indirecta
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// if err := db.Ping(); err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println("Conexión a la BD exitosa") //Sin el fmt y usando log indica la fecha y hora

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	handlers.ListContacts(db)

}