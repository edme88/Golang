package main

import (
	"fmt"
	"os"
)

func main(){
	//Pospone una ejecución al final
	defer fmt.Println(3)
	fmt.Println(1)
	fmt.Println(2)

	file, err := os.Create("hola.txt")
	if err!=nil{
		fmt.Println(err)
		return
	}

	defer file.Close()

	_, err = file.Write([]byte("Hola, Agus!"))
	if err != nil {
		fmt.Println(err)
		//file.Close()
		return
	}
	//file.Close()
}