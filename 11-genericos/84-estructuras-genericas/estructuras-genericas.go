package main

import (
	"fmt"
)

type Product[T uint | string] struct {
	//Id uint //BD relacional
	//Id string //si es una BD documental
	Id T
	Desc string
	Price float32
}

func main(){
	product1 := Product[uint]{1, "Zapatos", 50.12}
	product2 := Product[string]{"z01-m003", "Zapatos", 50.12}
	fmt.Println(product1)
	fmt.Println(product2)
}