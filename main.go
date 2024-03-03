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
	text = strings.TrimSpace(text)

	// Check if the argument is an existing .txt file
	if filepath.Ext(text) == ".txt" {
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

func EncodeArt(input string) string {
	lines := strings.Split(input, "\n")
	var result []string

	for _, line := range lines {
		encodedLine := encodeLineWithRunLength(line)
		result = append(result, encodedLine)
	}

	return strings.Join(result, "\n")
}

func encodeLineWithRunLength(line string) string {
	if len(line) == 0 {
		return ""
	}

	var result strings.Builder
	currentPattern := string(line[0])
	count := 1

	for i := 1; i < len(line); i++ {
		if line[i] == currentPattern[0] {
			count++
		} else {
			if count >= 5 {
				result.WriteString(fmt.Sprintf("[%d %s]", count, currentPattern))
			} else {
				result.WriteString(strings.Repeat(currentPattern, count))
			}
			currentPattern = string(line[i])
			count = 1
		}
	}

	// Append the last pattern after the loop
	if count >= 5 {
		result.WriteString(fmt.Sprintf("[%d %s]", count, currentPattern))
	} else {
		result.WriteString(strings.Repeat(currentPattern, count))
	}

	return result.String()
}
