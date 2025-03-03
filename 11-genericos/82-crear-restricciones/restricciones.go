package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

//En lugar de crear un constraints:
// type Numbers interface{
// 	~int | ~float64 | ~float32 | uint
// }

//En lugar de [T Numbers](nums ...T) T
func Sum[T constraints.Integer | constraints.Float](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}
	return total
}

func main(){
	fmt.Println(Sum(2,4.5))

	//Ir a https://pkg.go.dev/golang.org/x/exp/constraints
	//go mod init generic
	//go get golang.org/x/exp/constraints
	fmt.Println(Sum[uint](2,3,6,7))
	fmt.Println(Sum[float32](2.4,4.5))

}