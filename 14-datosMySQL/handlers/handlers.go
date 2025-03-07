package handlers

import (
	"database/sql"
	"datosMySQL/models"
	"fmt"
	"log"
)

//ListContacts lista todos los contactos desde la base de datos
func ListContacts(db *sql.DB) {
	//Consulta SQL para seleccionar todos los contactos
	query := "SELECT * FROM contact"

	//Ejecutar la consulta
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	//Iterar sobre los resultados y mostrarlos
	fmt.Println("\nLista de Contactos: ")
	fmt.Println("-----------------------")
	for rows.Next(){
		//Instancia de modelo contact
		contact := models.Contact{}

		var valueEmail sql.NullString
		err := rows.Scan(&contact.Id, &contact.Name, &valueEmail, &contact.Phone)
		if err != nil {
			log.Fatal(err)
		}

		if valueEmail.Valid{
			contact.Email = valueEmail.String
		}else{
			contact.Email = ""
		}

		fmt.Printf("ID: %d, Name: %s, Email: %s, Telefono: %s",
			contact.Id, contact.Name, contact.Email, contact.Phone)
			fmt.Println("-----------------------")
		}
}