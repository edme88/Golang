package main

import "fmt"

func main(){
	var i int

	for i<=10 {
		fmt.Println("El valor es: ", i)
		i++
	}

	for j := 1 ; j<=10 ; j++ {
		fmt.Println("El valor: ", j)

		if j==5 {
			//llego a 5, que salga
			break;
		}
	}

	for k := 1 ; k<=10 ; k++ {
		
		if k==5 {
			//llego a 5, que salte a la siguiente iteración (no imprime el 5)
			continue;
		}
		fmt.Println("El valor: ", k)

	}
}