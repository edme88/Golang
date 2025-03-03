package main

import "fmt"

func main(){
	saludo := hello("Agus")
	fmt.Println(saludo)
	sum, mul := calc(4,5)
	fmt.Println("La suma es: ", sum)
	fmt.Println("La multiplicación es: ", mul)
	suma, multi := calcu(4,5)
	fmt.Println("La suma es: ", suma)
	fmt.Println("La multiplicación es: ", multi)
}

//función que no devuelve nada
func hello0(name string) {
	fmt.Println("Hola "+name+" desde la función")
}

//Función que devuelve 1 valor
func hello(name string) string {
	return "Hola "+name+" desde la función"
}

//Funcion que devuelve 2 valores con tipo
func calc(a, b int) (int, int) {
	sum := a + b
	mul := a * b
	return sum, mul
}

//Funcion que devuelve 2 valores con nombre
func calcu(a, b int) (sum, mul int) {
	sum = a + b
	mul = a * b
	return
}