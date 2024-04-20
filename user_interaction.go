package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func getActionFromUser(reader *bufio.Reader) string {
	for {
		fmt.Println("Do you want to encode or decode? (e/d)")
		action, _ := reader.ReadString('\n')
		action = strings.TrimSpace(action)

		if action != "e" && action != "d" {
			fmt.Printf("%sInvalid option. Please choose 'e' for encode or 'd' for decode.%s\n", colorRed, colorReset)
			continue
		}
		return action
	}
}

func getInputTypeFromUser(reader *bufio.Reader) string {
	for {
		fmt.Println("Do you want to enter multiple lines or a single line? (m/s)")
		inputType, _ := reader.ReadString('\n')
		inputType = strings.TrimSpace(inputType)

		if inputType != "m" && inputType != "s" {
			fmt.Printf("%sInvalid option. Please choose 'm' for multiple lines or 's' for single line.%s\n", colorRed, colorReset)
			continue
		}
		return inputType
	}
}

func getTextArtFromUser(reader *bufio.Reader, inputType string) string {
	var text string
	if inputType == "m" {
		fmt.Println("Enter the text art (press Ctrl+D to finish):")
		var lines []string
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				break
			}
			lines = append(lines, line)
		}
		text = strings.Join(lines, "")
		text = strings.TrimSuffix(text, "\n")
	} else if inputType == "s" {
		fmt.Println("Enter the text art or filename:")
		text, _ = reader.ReadString('\n')
		text = strings.TrimSpace(text)
	}
	return text
}

func handleTextFile(text string) string {
	text = filepath.Join(directory, text)
	content, err := os.ReadFile(text)
	if err != nil {
		fmt.Printf("%sError reading file: %s%s\n", colorRed, err, colorReset)
		return ""
	}
	return string(content)
}

func encodeText(text string) {
	encodedText := EncodeArt(text)
	fmt.Println("Encoded text:")
	fmt.Println(encodedText)
}

func decodeText(text string) {
	decodedText, err := DecodeArt(text)
	if err != nil {
		fmt.Printf("%sError: %s%s\n", colorRed, err, colorReset)
		return
	}
	fmt.Println("Decoded text:")
	fmt.Println(decodedText)
}

func continueOperation(reader *bufio.Reader) bool {
	for {
		fmt.Println("Do you want to encode or decode another text? (y/n)")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)
		if choice != "y" && choice != "n" {
			fmt.Printf("%sInvalid option. Please choose 'y' or 'n'.%s\n", colorRed, colorReset)
			continue
		}
		return choice == "y"
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
