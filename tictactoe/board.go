package tictactoe

import ()

var board []byte
var playerMark, computerMark byte = 'X', 'O'
var winningMoves = [][]int{
	{0, 1, 2},
	{3, 4, 5},
	{6, 7, 8},
	{0, 3, 6},
	{1, 4, 7},
	{2, 5, 8},
	{0, 4, 8},
	{2, 4, 6},
}

func MakeBoard() []byte {
	board = make([]byte, 9)
	return board
}

func PopulateBoard(board []byte) {
	for cell := range board {
		board[cell] = '1' + byte(cell)
	}
}

func ThreeCellsMatch(list []int) bool {
	return board[list[0]] == board[list[1]] && board[list[1]] == board[list[2]]
}

func ComputerWon(list []int) bool {
	return board[list[0]] == computerMark
}

func GameIsWon() bool {
	for _, winningMove := range winningMoves {
		if ThreeCellsMatch(winningMove) {
			PrintBoard()
			if ComputerWon(winningMove) {
				PrintWinningMessage(computerMark)
			} else {
				PrintWinningMessage(playerMark)
			}
			return true
		}
	}
	return false
}

func OpenSquares() bool {
	for i, square := range board {
		if square == '1'+byte(i) {
			return true
		}
	}
	return false
}

func GameOver() bool {
	if GameIsWon() {
		return true
	} else if OpenSquares() {
		return false
	}
	PrintTieGame()
	return true
}
