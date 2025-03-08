package main

import (
	"bufio"
	"datosMySQL/database"
	"datosMySQL/handlers"
	"datosMySQL/models"
	"fmt"
	"os"
	"strings"

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

	// handlers.ListContacts(db)

	// handlers.GetContactByID(db, 3)

	//Crear una instancia de Contact
	// newContact := models.Contact{
	// 	Name: "Nuevo Usuario",
	// 	Email: "nuevo@example.com",
	// 	Phone: "123456789",
	// }
	//handlers.CreateContact(db, newContact)

	//ACTUALIZA una instancia de Contact
	// newContact := models.Contact{
	// 	Id: 6,
	// 	Name: "Usuario MUY Nuevo",
	// 	Email: "nuevoNuevo@example.com",
	// 	Phone: "987654321",
	// }

	// handlers.UpdateContact(db, newContact)

	// handlers.ListContacts(db)

	// handlers.DeleteContact(db, 6)

	// handlers.ListContacts(db)

	// var name, email, phone string
// 	fmt.Print("Ingrese un NOMBRE: ")
// 	fmt.Scanln(&name)
// 	fmt.Print("Ingrese un EMAIL: ")
// 	fmt.Scanln(&email)
// 	fmt.Print("Ingrese un TELEFONO: ")
// 	fmt.Scanln(&phone)

	for {
		fmt.Println(("\nMenú:"))
		fmt.Println("1. Listar contactos")
		fmt.Println("2. Obtener contacto por ID")
		fmt.Println("3. Crear nuevo contacto")
		fmt.Println("4. Actualizar contacto")
		fmt.Println("5. Eliminar contacto")
		fmt.Println("6. Salir")
		fmt.Print("Seleccione una opción: ")

		//Leer la opción seleccionada por el usuario
		var option int
		fmt.Scanln(&option)

		switch option {
		case 1:
			handlers.ListContacts(db)
		case 2:
			var contactId int
			fmt.Print("Ingrese un ID: ")
			fmt.Scan(&contactId)
			handlers.GetContactByID(db, contactId)
		case 3: 
			newContact := inputContactDetails(option)
			handlers.CreateContact(db, newContact)
			handlers.ListContacts(db)
		case 4:
			updateContact := inputContactDetails(option)
			handlers.UpdateContact(db, updateContact)
			handlers.ListContacts(db)
		case 5:
			var contactId int
			fmt.Print("Ingrese el ID del usuario que desea eliminar: ")
			fmt.Scan(&contactId)
			handlers.DeleteContact(db, contactId)
		case 6:
			fmt.Println("Saliendo del programa...")
			return
		default:
			fmt.Println("Opción incorrecta. Ingrese otra opción.")
		}
	}
}

//Función par aingresar los detalles del contacto desde la entrada estándar
func inputContactDetails(option int) models.Contact {
	//Leer la etrada del usuario utilizando bufio
	reader := bufio.NewReader(os.Stdin)

	var contact models.Contact

	if option == 4 {
		var contactId int
		fmt.Print("Ingrese un ID: ")
		fmt.Scan(&contactId)
		fmt.Scanln()
		contact.Id = contactId
	}

	fmt.Print("Ingrese el nombre del contacto: ")
	name, _ := reader.ReadString('\n')
	contact.Name = strings.TrimSpace(name)

	fmt.Print("Ingrese el email del contacto: ")
	email, _ := reader.ReadString('\n')
	contact.Email = strings.TrimSpace(email)

	fmt.Print("Ingrese el Teléfono del contacto: ")
	phone, _ := reader.ReadString('\n')
	contact.Phone = strings.TrimSpace(phone)

	return contact
}