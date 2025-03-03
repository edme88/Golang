package main

import "fmt"

// Recibe una función y devuelve una función
// permite programación modular y flexible re-utilizando funciones existentes
func double(f func(int) int, x int) int {
	return f(x * 2)
}

func addOne(x int) int{
	return x + 1
}

func main() {
	r := double(addOne, 3)

	fmt.Println("Resultado: ", r)
}