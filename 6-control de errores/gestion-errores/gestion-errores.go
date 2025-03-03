package main

import (
	"fmt"
	"strconv"
)

func divide(dividendo, divisor int)(int, error){
	if divisor == 0 {
		//return 0, errors.New("No se puede dividir en cero")
		return 0, fmt.Errorf("No se puededividir en CERO")
	}
	return dividendo/divisor, nil
}

func main(){
	//Manejar errores existentes
	str := "123f"
	strconv.Atoi(str)
	num, err := strconv.Atoi(str)

	if err != nil{
		fmt.Println("Error: ", err)
		//return
	}

	fmt.Println("Número: ", num)

	//Se pueden manejar errores personalizados
	fmt.Println(divide(3,2))
	resultado, error1 := divide(8,0)
	if error1!=nil{
		fmt.Println("Error: ", error1)
		return
	}
	fmt.Println(resultado)
}