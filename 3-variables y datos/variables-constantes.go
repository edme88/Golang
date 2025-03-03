package main

import (
	"fmt"

	"rsc.io/quote"
)

//Se recomieda que las constantes empiecen con mayúscula
const Pi float32 = 3.14
const E = 2.71828

const (
	X = 100
	Y = 0b1010 // binario
	Z = 0o12 // octal
	W = 0xFF //hexadecimal
)

const (
	Domingo = iota + 1
	Lunes
	Martes
	Miercoles
	Jueves
	Viernes
	Sabado
)

func main(){
	fmt.Println("Hola mundo!")
	fmt.Println(quote.Go())

	//Declaración de variables - post asignación (1)
	var name1, lastName1 string
	var age1 int

	name1 = "Agus"
	lastName1 = "Alici"
	age1 = 36

	// Declaración de variables "en bloque" (2)
	var (
		name2, lastName2 string
		age2 			 int
	)

	name2 = "Andres"
	lastName2 = "Leel"
	age2 = 36

	// Creación y asignación con tipo de datos (3)
	var (
		name3 string = "Pam"
		lastName3 string = "Alici"
		age3 int = 40
	)

	//Creación y Asignación (4)
	var (
		name4 = "Gui"
		lastName4 = "Alici"
		age4 = 38
	)

	//Creacion y Asignación inline (5)
	var name5, lastName5, age5 = "Bel", "Alici", 46

	//Las varibales pueden ser declaradas dentro o fuera de las funciones

	//creacion de variables sin var (solo dentro de funcionales)
	name6, lastName6, age6 := "Ari", "Edme", 7 //:= para crear nueva variable
	age6 = 8 //Solo para asignar

	// Constantes - Se suelen declarar a nivel de packete (fuera de la función)
	// También s epueden dentro de una función
	// Constantes se crean y asignan

	fmt.Println(name1, lastName1, age1)
	fmt.Println(name2, lastName2, age2)
	fmt.Println(name3, lastName3, age3)
	fmt.Println(name4, lastName4, age4)
	fmt.Println(name5, lastName5, age5)
	fmt.Println(name6, lastName6, age6)
	fmt.Println(Pi, E, X, Y, Z, W, Viernes)
}