package main

import (
	"datosMySQL/database"
	"datosMySQL/handlers"
	"datosMySQL/models"

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

	handlers.GetContactByID(db, 3)

	//Crear una instancia de Contact
	newContact := models.Contact{
		Name: "Nuevo Usuario",
		Email: "nuevo@example.com",
		Phone: "123456789",
	}

	handlers.CreateContact(db, newContact)

	// var name, email, phone string
// 	fmt.Print("Ingrese un NOMBRE: ")
// 	fmt.Scanln(&name)
// 	fmt.Print("Ingrese un EMAIL: ")
// 	fmt.Scanln(&email)
// 	fmt.Print("Ingrese un TELEFONO: ")
// 	fmt.Scanln(&phone)
}