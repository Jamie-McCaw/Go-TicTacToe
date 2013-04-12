package tictactoe

import (
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func TestComputerTurnFirstMove(t *testing.T) {
	PopulateBoard(MakeBoard())
	t.Log("Picks the best move")
	PlaceMove("1", playerMark)
	ComputerTurn()
	assert.Equal(t, board, OneMoveTestBoard())
	ClearBoard()
}

func TestComputerBlocksWin(t *testing.T) {
	t.Log("Blocks a winning move from a player")
	PlaceMove("1", playerMark)
	PlaceMove("2", playerMark)
	PlaceMove("5", computerMark)
	ComputerTurn()
	assert.Equal(t, board, MakeBlockingMove())
	ClearBoard()
}

func TestComputerTakesWin(t *testing.T) {
	t.Log("Takes a winning move")
	PlaceMove("1", computerMark)
	PlaceMove("2", computerMark)
	PlaceMove("4", playerMark)
	PlaceMove("5", playerMark)
	PlaceMove("8", playerMark)
	ComputerTurn()
	assert.Equal(t, board, MakesWinningMove())
	ClearBoard()
}

func TestComputerBlocksLaterInTheGame(t *testing.T) {
	t.Log("Blocks a later winning move")
	PlaceMove("2", playerMark)
	PlaceMove("8", computerMark)
	PlaceMove("4", playerMark)
	PlaceMove("3", computerMark)
	PlaceMove("6", playerMark)
	ComputerTurn()
	assert.Equal(t, board, MakesWinningMoveLateGame())
	ClearBoard()
}
