package tictactoe

import ()

func RunGame() {
	PrintHeader()
	PopulateBoard(MakeBoard())
	for {
		PlayerTurn()
		if GameOver() {
			break
		}
		ComputerTurn()
		if GameOver() {
			break
		}
	}
	PrintGameEnd()
}
