package tictactoe

import (
	"github.com/stretchrcom/testify/assert"
	"os"
	"testing"
)

func TestValidateInputForTrue(t *testing.T) {
	t.Log("Returns True for valid input")
	assert.True(t, ValidateInput("7"))
}

func TestValidateInputForFalse(t *testing.T) {
	t.Log("Returns False for an invalid input")
	assert.False(t, ValidateInput("b"))
}

func TestValidateInputForFalseAlso(t *testing.T) {
	t.Log("Returns False for invlaid number")
	assert.False(t, ValidateInput("10"))
}

func TestPlaceMovePlayer(t *testing.T) {
	t.Log("Places a move on the board")
	FakeBoard()
	mockboard[0] = playerMark
	PlaceMove("1", playerMark)
	assert.Equal(t, board, mockboard)
	ClearBoard()
}

func TestPlaceMoveComputer(t *testing.T) {
	t.Log("Places a move on the board")
	FakeBoard()
	mockboard[0] = computerMark
	PlaceMove("1", computerMark)
	assert.Equal(t, board, mockboard)
	ClearBoard()
}

func TestGetInput(t *testing.T) {
	read, write, _ := os.Pipe()
	oldStdin := os.Stdin
	os.Stdin = read
	defer func() {
		os.Stdin = oldStdin
	}()
	MockInput := "7"
	write.WriteString(MockInput)
	write.Close()
	t.Log("GetInput recieves 7 and passes it")
	assert.Equal(t, *GetInput(), "7")
}
