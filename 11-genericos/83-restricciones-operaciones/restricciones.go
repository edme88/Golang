package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func Includes[T comparable](list []T, value T) bool {
	for _, item := range list {
		if item == value {
			return true
		}
	}
	return false
}

func Filter[T constraints.Ordered](list []T, callback func(T) bool) []T {
	result := make([]T, 0, len(list))
	for _, item := range list {
		if callback(item) {
			result = append(result, item)
		}
	}
	return result
} 

func main(){
	stringsL := []string{"a", "b", "c", "d"}
	numbersL := []int{1,2,3,4,7,20,34,11}

	fmt.Println(Includes(stringsL, "a"))
	fmt.Println(Includes(stringsL, "f"))
	fmt.Println(Includes(numbersL, 4))
	fmt.Println(Includes(numbersL, 7))

	fmt.Println(Filter(numbersL, func(value int) bool {return value > 3}))
	fmt.Println(Filter(stringsL, func(value string) bool {return value >= "C"}))
}