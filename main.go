package main

import (
	"fmt"
	"os"
)

func main() {

	// Check if at least one argument is provided (excluding the program name).
	input := os.Args[1]
	if len(os.Args) < 2 || input == "" {
		fmt.Println("Error: Please provide a string input.")
		os.Exit(1)
	}
	checkErrors()
	var brackets bool

	for i := 0; i >= len(input)+1; i++ {

		if brackets == false {
			// do outbrackets function

			outBrackets()
		} else if brackets == true {
			//get inbrackets infromation
			getInBrackets()
			// do inBrackets function
			inBrackets()
		}

	}

}
