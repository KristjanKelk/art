package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var patternRE = regexp.MustCompile(`\[(.*?)\]`)
var directory = "Arts"

// ANSI color escape codes
const (
	colorReset = "\033[0m"
	colorRed   = "\033[31m"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		action := runOnServerOrTerminal(reader)
		switch action {
		case "t":
			handleTerminal(reader)
		case "s":
			startServer()
		default:
			fmt.Println("Invalid action. Please choose 't' for terminal or 's' for server.")
		}
	}
}

func handleTerminal(reader *bufio.Reader) {
	action := getActionFromUser(reader)
	if action == "" {
		return
	}

	inputType := getInputTypeFromUser(reader)
	if inputType == "" {
		return
	}

	text := getTextArtFromUser(reader, inputType)
	if text == "" {
		return
	}

	if filepath.Ext(strings.TrimSpace(text)) == ".txt" {
		text = handleTextFile(text)
	} else {
		// Handle other input types if needed
	}

	switch action {
	case "e":
		encodeText(text)
	case "d":
		decodeText(text)
	default:
		fmt.Println("Invalid action. Please choose 'e' for encode or 'd' for decode.")
	}

	if !continueOperation(reader) {
		os.Exit(0) // Exit the program if the user chooses not to continue
	}
}

func runOnServerOrTerminal(reader *bufio.Reader) string {
	fmt.Println("Do you want to run on a server or in a terminal? (s/t)")
	for {
		action, _ := reader.ReadString('\n')
		action = strings.TrimSpace(action)
		if action != "s" && action != "t" {
			fmt.Println("Invalid option. Please choose 's' for server or 't' for terminal.")
			continue
		}
		return action
	}
}
