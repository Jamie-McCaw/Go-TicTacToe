package tictactoe

import (
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func TestMakeBoard(t *testing.T) {
	t.Log("Returns a new blank board")
	assert.Equal(t, MakeBoard(), make([]byte, 9))
}

func TestPopulateBoard(t *testing.T) {
	t.Log("Returns a populate board")
	FakePopulatedBoard()
	assert.Equal(t, board, FakeBoard())
}

func TestComputerWonFalse(t *testing.T) {
	t.Log("Returns False for Computer Win")
	assert.False(t, ComputerWon(Row))
}

func TestComputerWonTrue(t *testing.T) {
	t.Log("Returns True for Computer Win")
	PutComputerInBoard()
	assert.True(t, ComputerWon(Row))
	ClearBoard()
}

func TestThreeCellsMatchFalse(t *testing.T) {
	t.Log("Returns false for 3 Cells Matching")
	assert.False(t, ThreeCellsMatch(Row))
}

func TestThreeCellsMatchTrue(t *testing.T) {
	t.Log("Returns true for 3 Cells Matching")
	PutComputerInBoard()
	assert.True(t, ThreeCellsMatch(Row))
	ClearBoard()
}

func TestOpenSquaresTrue(t *testing.T) {
	t.Log("Returns true if there are open squares")
	assert.True(t, OpenSquares())
}

func TestOpenSquaresFalse(t *testing.T) {
	t.Log("Returns false if there are no open squares")
	PutComputerInBoard()
	assert.False(t, OpenSquares())
	ClearBoard()
}

func TestGameIsWonByComputer(t *testing.T) {
	t.Log("Returns true if the game is won by the computer")
	PutComputerInBoard()
	assert.True(t, GameIsWon())
	ClearBoard()
}

func TestGameisWonByPlayer(t *testing.T) {
	t.Log("Returns true if the game is won by the player")
	PutPlayerInBoard()
	assert.True(t, GameIsWon())
	ClearBoard()
}

func TestGameIsWonFalse(t *testing.T) {
	t.Log("Returns false if the game is not won")
	PopulateNonWinningBoard()
	assert.False(t, GameIsWon())
	ClearBoard()
}

func TestGameOverForTie(t *testing.T) {
	t.Log("Returns true if the game is over because of a tie")
	TieGame()
	assert.True(t, GameOver())
	ClearBoard()
}

func TestGameOverForWin(t *testing.T) {
	t.Log("Returns true if the game is over because someone won")
	PutPlayerInBoard()
	assert.True(t, GameOver())
	ClearBoard()
}

func TestGameOverForFalse(t *testing.T) {
	t.Log("Returns false if the game is not over and there are open spaces")
	PopulateNonWinningBoard()
	assert.False(t, GameOver())
	ClearBoard()
}
