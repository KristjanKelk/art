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
	var resultArray []string
	length := len(input)

	// Splitting the input to slice/array for reworking into encoded version later.
	for i := 0; i < length; i++ {
		if i+1 < length && input[i] == input[i+1] {
			resultArray = append(resultArray, string(input[i]))
		} else if i+2 < length && input[i] == input[i+2] ||
			len(resultArray) > 0 && len(resultArray[len(resultArray)-1]) != 1 &&
				input[i:i+2] == resultArray[len(resultArray)-1] {
			resultArray = append(resultArray, input[i:i+2])
			i++
		} else {
			resultArray = append(resultArray, string(input[i]))
		}
	}

	// Constructing result string from the array by counting consecutive elements
	// and encoding if there are at least 5 consecutive elements
	var result string
	var count int
	i := 0
	length = len(resultArray)
	for i < len(resultArray) {
		if i+1 < length && resultArray[i] == resultArray[i+1] {
			count++
		} else {
			if count >= 4 {
				result += fmt.Sprintf(`[%d %s]`, count+1, resultArray[i])
				count = 0
			} else {
				result += resultArray[i]
				count = 0 // Reset count
			}
		}
		i++
	}
	return result
}
