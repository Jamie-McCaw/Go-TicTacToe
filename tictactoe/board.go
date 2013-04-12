package tictactoe

import (
	"strconv"
)

const boardSize = 9

var WinningPlayer byte
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
	board = make([]byte, boardSize)
	return board
}

func MoveAvailable(playermove string) bool {
	return board[playermove[0]-'1'] == playermove[0]
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

func WhoWon(winningMove []int) byte {
	if ComputerWon(winningMove) {
		return computerMark
	}
	return playerMark
}

func GameIsWon() bool {
	for _, winningMove := range winningMoves {
		if ThreeCellsMatch(winningMove) {
			WinningPlayer = WhoWon(winningMove)
			return true
		}
	}
	return false
}

func PlaceMove(move string, piece byte) {
	x, _ := strconv.Atoi(move)
	board[x-1] = piece
}

func OpenSquares() bool {
	for i, square := range board {
		if square == '1'+byte(i) {
			return true
		}
	}
	return false
}

func Tie() bool {
	if OpenSquares() {
		return false
	}
	return true
}

func GameOver() bool {
	if GameIsWon() || Tie() {
		return true
	}
	return false
}

func OpenMoves() []int {
	openSpaces := []int{}

	for index, space := range board {
		if SpaceIsNotTaken(space) {
			openSpaces = append(openSpaces, index)
		}
	}
	return openSpaces
}

func SpaceIsNotTaken(space byte) bool {
	return space != playerMark && space != computerMark
}
