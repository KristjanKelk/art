package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

var patternRE = regexp.MustCompile(`\[(.*?)\]`)
var directory = "Arts"

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Do you want to encode or decode? (e/d)")
	action, _ := reader.ReadString('\n')
	action = strings.TrimSpace(action)

	if action != "e" && action != "d" {
		fmt.Println("Invalid option. Please choose 'e' for encode or 'd' for decode.")
		return
	}

	fmt.Println("Enter the text art or file name:")
	text, _ := reader.ReadString('\n')
	//text = strings.TrimSpace(text)

	// Check if the argument is an existing .txt file
	if filepath.Ext(text) == ".txt" {
		text = strings.TrimSpace(text)
		text = filepath.Join(directory, text)
		content, err := os.ReadFile(text)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}
		text = string(content)
	}

	if action == "e" {
		encodedText := EncodeArt(text)
		fmt.Println("Encoded text:")
		fmt.Println(encodedText)
	} else if action == "d" {
		// If the argument is just a file name, prepend the directory name
		if !strings.Contains(text, "/") && !strings.Contains(text, "\\") {
			text = filepath.Join(directory, text)
		}

		// Check if the argument is an existing .txt file
		fileInfo, err := os.Stat(text)
		if err == nil && !fileInfo.IsDir() && filepath.Ext(text) == ".txt" {
			content, err := os.ReadFile(text)
			if err != nil {
				fmt.Println("Error reading file:", err)
				return
			}
			encodedText := string(content)
			decodedText, err := DecodeArt(encodedText)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Println("Decoded text:")
			fmt.Println(decodedText)
			return
		}

		// If the argument is not a file, treat it as encodedText directly
		decodedText, err := DecodeArt(text)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Decoded text:")
		fmt.Println(decodedText)
	}
}

func DecodeArt(encoded string) (string, error) {
	openBracketCount := strings.Count(encoded, "[")
	closeBracketCount := strings.Count(encoded, "]")

	// Check if square brackets marks are unbalanced
	if openBracketCount != closeBracketCount {
		return "", fmt.Errorf("square brackets are unbalanced")
	}

	matches := patternRE.FindAllStringSubmatchIndex(encoded, -1)

	var decoded strings.Builder
	lastEnd := 0
	for _, match := range matches {
		// Extract the substring within the current pair of brackets
		bracketSubstring := encoded[match[0]+1 : match[1]-1]

		// Iterate over the characters in the substring to find the separator
		var separatorIndex int
		foundSeparator := false
		for i, char := range bracketSubstring {
			if char == ' ' {
				separatorIndex = i
				foundSeparator = true
				break
			}
		}

		if !foundSeparator {
			return "", fmt.Errorf("space separator missing within brackets: %s", bracketSubstring)
		}

		// Extract countStr and repeatedStr based on the separator
		countStr := bracketSubstring[:separatorIndex]
		repeatedStr := bracketSubstring[separatorIndex+1:]

		// Check if countStr contains only digits
		if _, err := strconv.Atoi(countStr); err != nil {
			return "", fmt.Errorf("countStr must contain only digits")
		}

		// Append decoded string
		decoded.WriteString(encoded[lastEnd:match[0]])

		// Convert countStr to an integer
		count, err := strconv.Atoi(countStr)
		if err != nil {
			return "", fmt.Errorf("failed to convert count to integer: %w", err)
		}

		// Append repeatedStr count times
		for i := 0; i < count; i++ {
			decoded.WriteString(repeatedStr)
		}

		if repeatedStr == "" {
			return "", fmt.Errorf("No second argument")
		}

		lastEnd = match[1]
	}

	// Append any trailing text
	if lastEnd < len(encoded) {
		decoded.WriteString(encoded[lastEnd:])
	}

	return decoded.String(), nil
}

func EncodeArt(text string) string {

	var output strings.Builder
	count := 0
	currentPattern := ""

	for i := 0; i < len(text); i++ {
		if i < len(text)-1 && text[i] == text[i+1] {
			currentPattern = string(text[i])
			count++
			continue
		} else {
			if len(currentPattern) != 1 && i < len(text)-3 && text[i:i+2] == text[i+2:i+4] {
				currentPattern = text[i : i+2]
				count++
				i++
				continue
			}
		}

		if count > 0 {
			if count >= 2 {
				fmt.Fprintf(&output, "[%d %s]", count+1, string(currentPattern))
			} else {
				fmt.Fprint(&output, RepeatString(count+1, string(currentPattern)))
			}
			count = 0
			if len(currentPattern) == 2 {
				i++
			}
			currentPattern = ""
		} else {
			fmt.Fprint(&output, string(text[i]))
		}
	}
	return output.String()
}

func RepeatString(n int, text string) string {
	return strings.Repeat(text, n)
}
