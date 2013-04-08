package ttt

import (
	"fmt"
	"regexp"
)

func OutputBoard() {
	fmt.Printf(" %v | %v | %v \n", 0, 1, 2)
	fmt.Println("-----------")
	fmt.Printf(" %v | %v | %v \n", 3, 4, 5)
	fmt.Println("-----------")	
	fmt.Printf(" %v | %v | %v \n", 6, 7, 8) 
}

func OutputPrompt() {
	fmt.Print("-> ")
	userInput := GetInput()
	fmt.Println(userInput)
}

func GetInput() string {
	var input string
	fmt.Scan(&input)
	if ValidateInput(input) == true {
		return input
	}
	GetInput()
	return input
}

func OutputGreeting() {
	fmt.Println("Welcome To Tic Tac Toe\n")
}

func ValidateInput(input string) bool {
	inputMatch, _ := regexp.MatchString("^[0-8]$", input)
	return inputMatch
}




