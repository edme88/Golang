package main

import (
	"fmt"
	"strconv"

	"rsc.io/quote"
)

func main(){
	//Conversion de tipos de datos
	fmt.Println("Hola Mundo")
	fmt.Println(quote.Go())

	var integer16 int16 = 50
	var integer32 int32 = 100

	fmt.Println(int32(integer16) + integer32)

	s := "100s"
	i, _ := strconv.Atoi(s)
	fmt.Println(i)

	n := 42
	s = strconv.Itoa(n)
	fmt.Println(s)
	fmt.Println(s+s)

	//Info del packete fmt https://pkg.go.dev/fmt

	fmt.Print("Chau!!!")
	fmt.Print("linea pegada")

	name := "Agus"
	age := 36

	fmt.Printf("Hola, me llamo %s y tengo %d años. \n", name, age)

	greeting := fmt.Sprintf("Hola, me llamo %s y tengo %d años.", name, age)

	fmt.Println(greeting)

	var nameUser, name2 string
	var yourAge int

	fmt.Print("Ingrese sus nombres: ")
	fmt.Scanln(&nameUser, &name2)
	fmt.Print("Ingrese su edad: ")
	fmt.Scanln(&yourAge)
	fmt.Printf("Hola, me llamo %s %s y tengo %d años. \n", nameUser, name2, yourAge)

	//Se pude ysar %v si no conocemos el tipo de datos
	fmt.Printf("El tipo de name es: %T \n", name)
}