package main

import (
	"fmt"
	"math/rand"
)

func main(){
	play()
}

func play(){
	numAleatorio := rand.Intn(100)
	var numIngresado int
	var cantIntentos int
	const maxIntentos = 10

	for cantIntentos < maxIntentos {
		cantIntentos++
		fmt.Printf("Ingrese un número (intentos restantes: %d): ", maxIntentos - cantIntentos + 1)
		fmt.Scanln(&numIngresado)

		if numIngresado == numAleatorio {
			fmt.Println("¡Felicitaciones, adivinaste el número!")
			playAgain()
			return //para indicar que termina la ejecución
		} else if (numIngresado < numAleatorio) {
			fmt.Println("No acertaste. El número es mayor")
		} else{
			fmt.Println("No acertaste. El número es menor.")
		}
	}
	fmt.Println("Se acabaron los intentos. El número era: ", numAleatorio)
	playAgain()
}

func playAgain(){
	var option string
	fmt.Println("Desea jugar nuevamente? (s/n): ")
	fmt.Scanln(&option)

	switch option{
	case "s":
		play()
	case "n":
		fmt.Println("¡Gracias por jugar!")
	default:
		fmt.Println("Opción incorrecta. Intente de nuevo")
		playAgain()
	}
}