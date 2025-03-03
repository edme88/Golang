package main

import "fmt"

// Recuerda la definición de una función exterior
func incrementer() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := incrementer()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
}