package tictactoe

import (
	"github.com/stretchrcom/testify/assert"
	"os"
	"testing"
)

func TestPlayerTurnGoodMove(t *testing.T) {
	read, write, _ := os.Pipe()
	oldStdin := os.Stdin
	os.Stdin = read
	defer func() {
		os.Stdin = oldStdin
	}()
	MockInput := "6"
	write.WriteString(MockInput)
	write.Close()
	t.Log("Start the PlayersTurn and submit good value that gets placed on the board")
	PlayerTurn()
	assert.Equal(t, board, OneMoveTestBoard())
}
