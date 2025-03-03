package main

import "fmt"

func factorial(n int) int{
	if n == 0{
		return 1
	}
	return n * factorial(n-1)
}

//Ejemplo de factorial
// factorial !5 = 1 * 2 * 3 * 4 *5
func main(){
	fmt.Println(factorial(5))
}