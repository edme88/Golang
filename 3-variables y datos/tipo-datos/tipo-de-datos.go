package main

import (
	"fmt"
	"math"
)

func main(){
	//Tipos de datos básicos: 4
	//1. numeros
	//2. cadena de caracteres
	//3. boleanos
	var miNum int8 = 127
	//int8 almacena 8 bits,siendo de -128 hasta 127
	//si se coloca solo int permite almacenar entereos de 32 o 64 acorde a nuestro SO
	//el uint numeros enteros positivos

	var miNumComoso float32 = 23.23
	//Float de 32 o 64

	fmt.Println(miNum, miNumComoso)
	fmt.Println(math.MinInt64,math.MaxInt64)
	fmt.Println(math.MaxFloat32, math.MaxFloat64, math.MaxUint8)

	var miBooleano bool = true

	fullName := "Agus Ali \" o tambien \t(alias) y \n"

	fmt.Println(miBooleano, fullName)
	//tipo de datos como byte
	var a byte = 'a';
	fmt.Println(a) //Imprime el 97 que es el valor en codigo ascii

	s := "hola"
	fmt.Println(s[0]) //104 valor de h en decimal

	//tipo de datos rune - carcateres unicode
	var r rune = '❤'
	fmt.Println(r) //valor unicode

	//Tipos agregados
	//1. matrices
	//2. estructuras de datos - colecciones de datos

	//De referencia
	//1. punteros
	//2. segmentos
	//3. mapas
	//4. funciones
	//5. canales

	//De interface

	//Valor predeterminado de las variables
	//Ej. un booleano arranca como falso
	//Ej. String arranca como cadena vacio

	var (
		defaultInt int
		defaultUint uint
		defaultFloat float32
		defaultBool bool
		defaultString string
	)

	fmt.Println(defaultInt, defaultUint, defaultFloat, defaultBool, defaultString)
}