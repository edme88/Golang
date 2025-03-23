package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//usuario:password@tcp(ruta:puerto)/database
const url =  "root:GoLang123#@tcp(localhost:3306)/goweb_db"
var db *sql.DB

//Realiza la Conexión
func Connect(){
	conection, err := sql.Open("mysql", url)

	if err != nil {
		panic(err)
	}

	db = conection
	fmt.Println("Conexión exitosa")
}

//Cierra la conexión
func Close(){
	db.Close()
}

//Verificar la conexión
func Ping(){
	if err := db.Ping(); err != nil {
		panic(err)
	}
}

//Crea una tabla en la base de datos
func CreateTable(schema string, name string){
	if !ExistsTable(name){
		_, err := db.Exec(schema)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}

//Verifica si una tabla existe o no
func ExistsTable(tableName string) bool{
	sql := fmt.Sprintf("SHOW TABLES LIKE '%s'", tableName)
	rows, err := db.Query(sql)
	if err != nil{
		fmt.Println("Error:", err)
	}
	return rows.Next()
}