package tictactoe

import ()

func TestRunGame(t *testing.T) {
	read, write, _ := os.Pipe()
	oldStdin := os.Stdin
	os.Stdin = read
	defer func() {
		os.Stdin = oldStdin
	}()
	testString := "6"
	write.WriteString(testString)
	write.Close()
	t.Log("Start the PlayersTurn and submit good value that gets placed on the board")
	PlayerTurn()
	assert.Equal(t, board, OneMoveTestBoard())
}