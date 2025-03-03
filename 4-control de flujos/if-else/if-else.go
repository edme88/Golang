package main

import (
	"fmt"
	//"math"
	"time"
)

func main(){

	// t := time.Now()
	// hora := t.Hour()

	// //if no lleva parentesis
	// if hora<12{
	// 	fmt.Println("Mañana")
	// } else if hora < 18{
	// 	fmt.Println("Tarde")
	// } else{
	// 	fmt.Println("Noche")
	// }

	if t := time.Now(); t.Hour() < 12{
		fmt.Println("Mañana")
	} else if t.Hour() < 18{
		fmt.Println("Tarde")
	} else{
		fmt.Println("Noche")
	}
}