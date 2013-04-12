package tictactoe

import ()

func PlayerTurn() {
	PrintBoard()
	PrintPrompt()
	playerMove := GetInput()
	if !ValidateInput(*playerMove) {
		PrintInvalidMove()
		PlayerTurn()
	} else {
		if MoveAvailable(*playerMove) {
			PlaceMove(*playerMove, playerMark)
		} else {
			PrintInvalidMove()
			PlayerTurn()
		}
	}

}
