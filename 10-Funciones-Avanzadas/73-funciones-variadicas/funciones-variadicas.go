package main

import "fmt"

func suma(nums ...int) int {
	//fmt.Printf("%T - %v\n", nums, nums) imprime []int - [12 45 78 56]

	var total int
	for _, num := range nums {
		total += num
	}
	return total
}

func sumaVariada(name string, nums ...int) int {
	//fmt.Printf("%T - %v\n", nums, nums) imprime []int - [12 45 78 56]

	var total int
	for _, num := range nums {
		total += num
	}
	fmt.Printf("Hola %s! La suma es %d\n", name, total)
	return total
}

//recibir n datos de n tipos
func imprimirDatos(datos ...interface{}){
	for _, dato := range datos {
		fmt.Printf("%T - %v\n", dato, dato)
	}
}

func main(){
	sumaVariada("Agus", 12,34,56);
	fmt.Println("La suma es: ", suma(12,45,78,56))

	imprimirDatos("hola", 3, "chau", true, 2.34)
}