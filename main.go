package main

import (
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
	if len(os.Args) != 2 {
		fmt.Println("Error: Please provide the filename or encoded text as an argument")
		return
	}

	arg := os.Args[1]

	// If the argument is just a file name, prepend the directory name
	if !strings.Contains(arg, "/") && !strings.Contains(arg, "\\") {
		arg = filepath.Join(directory, arg)
	}

	// Check if the argument is an existing .txt file
	fileInfo, err := os.Stat(arg)
	if err == nil && !fileInfo.IsDir() && filepath.Ext(arg) == ".txt" {
		content, err := os.ReadFile(arg)
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
		fmt.Println(decodedText)
		return
	}

	// If the argument is not a file, treat it as encodedText directly
	decodedText, err := DecodeArt(arg)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(decodedText)
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
