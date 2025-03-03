package main

import (
	"log"
	"os"
)

func main(){
	log.Print("Registro")
	log.Println("otro registro")
	//muestra
	//2025/03/02 11:59:09 Registro
	//2025/03/02 11:59:09 otro registro
	//log.Fatal("Esto detendra el programa")
	//log.Panic("Es un panico logueado")
	log.SetPrefix("El error que se genera en MAIN es:")
	log.Print("Registro")
	log.Println("otro registro")
	//log.Fatal("Esto detendra el programa")
	log.Print("Se ve algo?")

	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644) //0644 es permiso de lectura y escritura para el propietario y para la aplicación
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.SetOutput(file)
	log.Print("Es un log")
}