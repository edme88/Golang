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

		fmt.Printf("ID: %d, Name: %s, Email: %s, Telefono: %s\n",
			contact.Id, contact.Name, contact.Email, contact.Phone)
			fmt.Println("-----------------------")
		}
}


//GetContactByID obitnt un contacto de la base de datos mediante su ID
func GetContactByID(db *sql.DB, contactID int){
	query := "SELECT * FROM contact WHERE id = ?"
	row := db.QueryRow(query, contactID)

	//Instancia de modelo contact
	contact := models.Contact{}
	var valueEmail sql.NullString //Para manejar el valor null

	//Escanear el resultado en el modelo contact
	err := row.Scan(&contact.Id, &contact.Name, &valueEmail, &contact.Phone)
	if err != nil {
		if err == sql.ErrNoRows{
			log.Fatalf("No se encontró ningpun contacto con el ID %d", contactID)
		}
		log.Fatal(err)
	}

	if valueEmail.Valid{
		contact.Email = valueEmail.String
	}else{
		contact.Email = ""
	}

	fmt.Println("-----------------------")
	fmt.Println("CONTACTO BUSCADO: ")
	fmt.Printf("ID: %d, Name: %s, Email: %s, Telefono: %s \n",
		contact.Id, contact.Name, contact.Email, contact.Phone)
		fmt.Println("-----------------------")
}

//CreateContact registra un nuevo contacto rn la base de datos
func CreateContact(db *sql.DB, contact models.Contact){
	

	query := "INSERT INTO contact (name, email, phone) values (?, ?, ?);";
	_, err := db.Exec(query, contact.Name, contact.Email, contact.Phone)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Nuevo contacto registrado con éxito")
}

// UpdateContact actualiza un contacto existente en la base de datos
func UpdateContact(db *sql.DB, contact models.Contact){
	query := "UPDATE contact SET name = ?, email  = ?, phone = ? WHERE id = ?";

	_, err := db.Exec(query, contact.Name, contact.Email, contact.Phone, contact.Id)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Contacto actualizado con éxito")
}

// DeleteContact elimina un contacto de la base de datos
func DeleteContact(db *sql.DB, contactID int){
	query := "DELETE FROM contact where Id = ?";
	_, err := db.Exec(query, contactID)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Contacto eliminado con éxito")
}