package rps

import (
	"math/rand"
	"strconv"
)

const (
	ROCK     = 0 // Piedra vence a tijeras (tijeras + 1) %  3   =  0
	PAPER    = 1 // Papel vence a piedra 	 (piedra  + 1) %  3   =  1
	SCISSORS = 2 // Tijeras vencen a papel (papel   + 1) %  3   =  2
)

type Round struct {
	Message           string `json:"message"`
	ComputerChoice    string `json:"computer_choice"`
	RoundResult       string `json:"round_result"`
	ComputerChoiceInt int    `json:"computer_choice_int"`
	ComputerScore     string `json:"computer_score"`
	PlayerScore       string `json:"player_score"`
}

// Mensajes ganadores
var winMessages = []string{
	"¡Bien hecho!",
	"¡Buen trabajo!",
	"Deberias probar tu suerte en la loteria",
}

// Mensajes perdedores
var loseMessages = []string{
	"¡Que lastima!",
	"¡Intentalo de nuevo!",
	"Hoy no es tu dia amigo",
}

// Mensaje empate
var drawMessages = []string{
	"Las grandes mentes piensan igual",
	"Oh oh. Intentalo de nuevo",
	"Nadie gana, pero puedes intentarlo de nuevo",
}

// Variables para el puntaje
var ComputerScore, PlayerScore int

func PlayRound(playerValue int) Round {
	var computerChoice, roundResult string
	var computerChoiceInt int

	// Elecciones de PC
	computerValue := rand.Intn(3)
	switch computerValue {
	case ROCK:
		computerChoiceInt = ROCK
		computerChoice = "La computadora eligio PIEDRA"
	case PAPER:
		computerChoiceInt = PAPER
		computerChoice = "La computadora eligio PAPEL"
	case SCISSORS:
		computerChoiceInt = SCISSORS
		computerChoice = "La computadora eligio TIJERAS"
	}
	// variable para contener mensaje
	var message string
	// generar mensaje aleatorio
	messageInt := rand.Intn(3)
	if playerValue == computerValue {
		roundResult = "Empate"
		message = drawMessages[messageInt]
	} else if playerValue == (computerValue+1)%3 {
		PlayerScore++
		roundResult = "¡Jugador Gana!"
		message = winMessages[messageInt]
	} else {
		ComputerScore++
		roundResult = "¡PC Gana!"
		message = loseMessages[messageInt]
	}

	return Round{
		Message:           message,
		ComputerChoice:    computerChoice,
		RoundResult:       roundResult,
		ComputerChoiceInt: computerChoiceInt,
		ComputerScore:     strconv.Itoa(ComputerScore), // convertimos int a string
		PlayerScore:       strconv.Itoa(PlayerScore),
	}
}
