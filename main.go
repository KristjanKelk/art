package main

import (
	"fmt"
	"os"
)

func main() {
	// Check if at least one argument is provided (excluding the program name).
	if len(os.Args) < 2 || os.Args[1] == "" {
		fmt.Println("Error: Please provide a string input.")
		os.Exit(1)
	}

	// Takes input string into userInput
	userInput := os.Args[1]
	fmt.Println("User input:", userInput)
	result := ""
	//decodeing input string
	for i, char := range userInput {
		var inBrackets bool
		if char == '[' {
			if rune(userInput[i+1]) < '0' || rune(userInput[i+1]) > '9' {
				fmt.Println("Error: The first argument is not a number.")
				os.Exit(1)
			}
			inBrackets = true
			continue
		}
		if char == ']' {
			if inBrackets == false {
				fmt.Println("Error: Bracket error.")
				os.Exit(1)
			} else {
				inBrackets = false
				continue
			}
		}
		result += string(char)
	}
	println(result)
}
