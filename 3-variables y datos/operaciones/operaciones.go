package main

import (
	"fmt"
	"math"
)

func main(){
	a := 10.0
	b := 3.0

	c := 10
	d := 3

	c++
	b--
	a = a + 5
	a += 5
	a -= 5
	fmt.Println(a)
	fmt.Println(a + b)
	fmt.Println(a / b)
	fmt.Println(c % d)
	fmt.Println(math.Pi)
	fmt.Println(math.E)
	fmt.Println(math.Pow(2,3))
}