package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var patternRE = regexp.MustCompile(`\[(\d+) ([^\]]+)\]`)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error")
		return
	}

	encodedText := os.Args[1]
	decodedText, err := DecodeArt(encodedText)
	if err != nil {
		fmt.Println("Error")
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
		countStr := encoded[match[2]:match[3]]
		repeatedStr := encoded[match[4]:match[5]]
		// Check if the second argument is an empty string
		if repeatedStr == "" {
			return "", fmt.Errorf("second argument is an empty string")
		}

		// Check if countStr is not a number
		count, err := strconv.Atoi(countStr)
		if err != nil {
			return "", fmt.Errorf("first argument is not a number")
		}

		// Check if the arguments are not separated by a space
		if match[3] != match[4]-1 {
			return "", fmt.Errorf("arguments are not separated by a space")
		}

		// Append decoded string
		decoded.WriteString(encoded[lastEnd:match[0]])
		decoded.WriteString(strings.Repeat(repeatedStr, count))

		lastEnd = match[1]
	}

	// Append any trailing text
	if lastEnd < len(encoded) {
		decoded.WriteString(encoded[lastEnd:])
	}

	return decoded.String(), nil
}
