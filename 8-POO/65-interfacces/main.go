package main

import (
	"animals-char/animal"
)

func main() {
	miPerro := animal.Perro{Nombre: "Laica"}
	miGato := animal.Gato{Nombre: "Tom"}

	//permite flexibilidad y reutilizacion de código
	animal.HacerSonido(&miPerro)
	animal.HacerSonido(&miGato)

	animales := []animal.Animal {
		&animal.Perro{Nombre: "Fido"},
		&animal.Gato{Nombre: "Eustaquio"},
		&animal.Perro{Nombre: "Buddy"},
		&animal.Gato{Nombre: "Luna"},
	}

	for _, animal := range animales {
		animal.Sonido()
	}
}