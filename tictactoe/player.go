package tictactoe

import (
	"strconv"
)

func PlayerTurn() {
	PrintBoard()
	PrintPrompt()
	playerMove := GetInput()
	if !ValidateInput(*playerMove) {
		PrintInvalidMove()
		PlayerTurn()
	} else {
		if MoveAvailable(*playerMove) {
			x, _ := strconv.Atoi(*playerMove)
			board[x-1] = playerMark
		} else {
			PrintInvalidMove()
			PlayerTurn()
		}
	}

}
