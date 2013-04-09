package tictactoe

import (
	"fmt"
	"bufio"
	"os"
)

var input = bufio.NewReader(os.Stdin)

func PrintBoard() {
    fmt.Printf("%s\n%s\n%s\n", board[0:3], board[3:6], board[6:9])
}

func PrintPrompt() {
	fmt.Print("-> ")
}

func PrintHeader() {
	fmt.Println("Enter a Number")
}

func PrintTieGame() {
	fmt.Println("Tie Game.")
}

func PrintWinningMessage(mark byte) {
	fmt.Printf("%v Wins!\n", string(mark))
}
