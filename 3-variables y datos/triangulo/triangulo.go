package main

import (
	"fmt"
	"math"
)

func main(){
	var lado1, lado2 float64
	const precision = 3

	//Entrada de Datos
	fmt.Print("Ingrese lado 1: ")
	fmt.Scanln(&lado1)
	fmt.Print("Ingrese lado 2: ")
	fmt.Scanln(&lado2)

	// Calculos
	var hipotenusa float64 = math.Sqrt(math.Pow(lado1,2) + math.Pow(lado2,2))
	area := (lado1 * lado2)/2
	perimetro := (lado1 + lado2 + hipotenusa)

	//Salida de Datos
	fmt.Printf("Área: %.2f \n", area) //para mostrar 2 decimales
	fmt.Printf("Perímetro: %.2f \n", perimetro) //para mostrar 2 decimales

	fmt.Printf("Área: %.*f \n", precision, area)
	fmt.Printf("Perímetro: %.*f \n", precision, perimetro)
}