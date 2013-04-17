package tictactoe

import ()

import (
	"github.com/stretchrcom/testify/assert"
	"os"
	"testing"
)

func TestRunGame(t *testing.T) {
	RunGameTestBoard()
	read, write, _ := os.Pipe()
	oldStdin := os.Stdin
	os.Stdin = read
	defer func() {
		os.Stdin = oldStdin
	}()
	MockInput := "7"
	write.WriteString(MockInput)
	write.Close()
	t.Log("Runs through RunGame and finishes the game")
	PlayerTurn()
	assert.Equal(t, board, RunGameMockBoard())
	ClearBoard()
}
