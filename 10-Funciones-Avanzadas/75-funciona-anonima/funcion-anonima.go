package main

import "fmt"

func saludar(name string, f func(string)){
	f(name)
}

func duplicar(n int) int {
	return n * 2
}

func triplicar(n int) int {
	return n * 3
}

func aplicar(f func(int) int, n int) int {
	return f(n)
}

func main() {
	func () {
		fmt.Println("Hola soy una función anonima")
	}() //sin los parentesis al final no se ejecuta

	saludo := func (name string) {
		fmt.Printf("Hola %s!", name)
	}

	saludo("Agus")

	saludar("Andres", saludo)

	r1 := aplicar(duplicar, 5)
	r2 := aplicar(duplicar, 6)

	fmt.Println(r1, r2)
}