package main

import "fmt"

type integer int  //forma de crear un tipo de datos nuevo

// [T int] se llama restricción arbitrario
// [T int | float64] restricción de unión de elementos
// [T ~int | ~float64] restricciones de aproximación tienen en cuenta el nuevo tipo de datos creado
func Sum[T ~int | ~float64](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}
	return total
}

func main(){
	var num1 integer = 100
	var num2 integer = 20
	fmt.Println(Sum(1,5,8))
	fmt.Println(Sum[int](1,5,8))
	fmt.Println(Sum(1.2,5.4,8.3))
	fmt.Println(Sum[float64](1.2,5.4,8.3))
	fmt.Println(Sum(1,2,2.3))
	fmt.Println(Sum(num1,num2))
}