package main

import "fmt"

func PrintList(list ...interface{}){
	for _,value := range list {
		fmt.Println((value))
	}
}

//Funcion variadica
func PrintListNueva(list ...any){
	for _,value := range list {
		fmt.Println((value))
	}
}

func main(){

	PrintList("Agus", 23.45, true, 12)
	PrintListNueva("Agus", 23.45, true, 12)
}