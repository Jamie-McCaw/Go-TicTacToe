package tictactoe

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

func PrintInvalidMove() {
	fmt.Println("Invalid Move Choice\n")
}

func ValidateInput(input string) bool {
	inputMatch, _ := regexp.MatchString("^[1-9]$", input)
	return inputMatch
}

func GetInput() *string {
	var input string
	fmt.Scan(&input)
	return &input
}
