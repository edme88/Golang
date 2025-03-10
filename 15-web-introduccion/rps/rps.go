package rps

import (
	"math/rand"
	"strconv"
)

const (
	ROCK = 0 //Piedra. Vence a las tijeras. (tijeras + 1) % 3 = 0
	PAPER = 1 //Papel. Vence a la piedra. (piedra + 1) % 3 = 1
	SCISSORS = 2 // Tijera. Vence a papel (papel + 1) % 3 = 2
)

//Estructura para dar resultado de cada ronda
type Round struct {
	Message string `json:"message"`
	ComputerChoice string `json:"computer_choice"`
	RoundResult string `json:"round_result"`
	ComputerChoiceInt int `json:"computer_choice_int"`
	ComputerScore string `json:"computer_score"`
	PlayerScore string `json:"player_score"`
}

//Mensajes para cuando gana
var winMessages = []string{
	"Bien echo!",
	"Felicitaciones, Ganaste!",
	"Has ganado esta ronda!",
}

//Mensajes para cuando pierde
var loseMessages = []string{
	"Qué lástima!",
	"Inténtalo de nuevo!",
	"Hoy no es tu día!",
}

//Mensajes de empate
var drawMessages = []string{
	"Las grande smentes piensan igual.",
	"Nadie gana, pero puedes intentarlo de nuevo.",
	"Parece que seleccionamos lo mismo!",
}

//Variables para el puntaje
var ComputerScore, PlayerScore int

func PlayRound(playerValue int) Round {
	computerValue := rand.Intn(3) //Número aleatorio

	var computerChoice, roundResult string
	var computerChoiceInt int 

	//Mensaje dependiendo de lo que elegio la computadora
	switch computerValue {
	case ROCK:
		computerChoiceInt = ROCK
		computerChoice = "La computadora eligió PIEDRA"
	case PAPER:
		computerChoiceInt = PAPER
		computerChoice = "La computadora eligió PAPEL"
	case SCISSORS:
		computerChoiceInt = SCISSORS
		computerChoice = "La computadora eligió TIJERAS"
	}

	//Generar número aleatorio de 0-2
	messageInt := rand.Intn(3)


	//declarar una var para contener el mesaje
	var message string

	if playerValue == computerValue {
		roundResult = "Es un empate"
		//Seleccionar mensaje de drawMessages
		message = drawMessages[messageInt]
	} else if playerValue == (computerValue+1)%3 {
		PlayerScore++
		roundResult = "El jugador GANA!!!"
		//Seleccionar mensaje de drawMessages
		message = winMessages[messageInt]
	} else {
		ComputerScore++
		roundResult = "La Computadora GANA!!!"
		//Seleccionar mensaje de drawMessages
		message = loseMessages[messageInt]
	}

	//Retornar el resultado
	return Round{
		Message: message,
		ComputerChoice: computerChoice,
		RoundResult: roundResult,
		ComputerChoiceInt: computerChoiceInt,
		ComputerScore: strconv.Itoa(ComputerScore),
		PlayerScore: strconv.Itoa(PlayerScore),
	}
}