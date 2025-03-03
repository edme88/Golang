package main

import (
	"fmt"
	"runtime"
	"time"
)

func main(){
	// os := runtime.GOOS

	// //No es necesario la instrucción Break en cada case
	// switch os{
	// case "windows":
	// 	fmt.Println("Go run -> Windows")
	// case "linux":
	// 	fmt.Println("Go run -> Linux")
	// case "darwin":
	// 	fmt.Println("Go run -> macOs")
	// default:
	// 	fmt.Println("Go run -> en otro OS")
	// }

	switch t := time.Now(); {
	case t.Hour() < 12:
		fmt.Println("Mañana")
	case t.Hour() < 18:
		fmt.Println("Tarde")
	default:
		fmt.Println(t.Hour(), "Noche")
	}

	switch runtime.GOOS{
	case "windows":
		fmt.Println("Go run -> Windows")
	case "linux":
		fmt.Println("Go run -> Linux")
	case "darwin":
		fmt.Println("Go run -> macOs")
	default:
		fmt.Println("Go run -> en otro OS")
	}
}