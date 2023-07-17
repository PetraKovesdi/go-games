package rps

import (
	"math/rand"
	"time"
)

type Round struct {
	ComputerChoice string `json:"computer_choice"`
	Result         string `json:"round_result"`
	Score          int    `json:"score"`
	Message        string `json:"message"`
}

var player_score = 0

func PlayRound(playerValue int) Round {
	options := []string{"ROCK", "PAPER", "SCISSORS"}
	rand.Seed(time.Now().UnixNano())
	computerValue := rand.Intn(3)
	computerChoice := options[computerValue]

	var result string

	if playerValue == computerValue {
		result = "It is a draw."
	} else if playerValue == (computerValue+1)%len(options) {
		result = "Player wins."
		player_score += 1
	} else {
		result = "Computer wins."
		player_score -= 1
	}

	var RoundResult Round

	RoundResult.ComputerChoice = computerChoice
	RoundResult.Result = result
	RoundResult.Score = player_score
	if player_score > 3 {
		RoundResult.Message = "Stanley is proud of you."
	} else {
		RoundResult.Message = ""
	}

	return RoundResult

}
