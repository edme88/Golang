package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/joho/godotenv"
)

func Connect() (*sql.DB, error) {
	//Cargar los valores desde las variables de entorno
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	dns := fmt.Sprintf("%s:%v@tcp(%s:%v)/%v", 
	os.Getenv("DB_USER"),
	os.Getenv("DB_PASSWORD"),
	os.Getenv("DB_HOST"),
	os.Getenv("DB_PORT"),
	os.Getenv("DB_NAME")) //datos de la BD para poder conectarnos usuario:Password@tcp(localhost:port)
	
	//Abrir la conexión con la BD
	db, err := sql.Open("mysql", dns) //Se utiliza el pacquete de forma indirecta
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Conexión a la BD exitosa") //Sin el fmt y usando log indica la fecha y hora
	return db, nil
}