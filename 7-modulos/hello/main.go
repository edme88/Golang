package main

import (
	"fmt"
	"log"

	"github.com/edme88/greetings"
)

func main() {
	log.SetPrefix("Greetings: ")
	log.SetFlags(0) //Bandera de formato - Determina que información adicional se coloca en los registros
	
	names := []string{"Agus","Edme","Andres"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
	
	message, err := greetings.Hello("Agus")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}