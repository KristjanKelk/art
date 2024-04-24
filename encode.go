package main

import (
	"fmt"
	"strings"
)

func EncodeArt(text string) string {
	var output strings.Builder
	count := 0
	currentPattern := ""

	// Iterate through each character in the input text.
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

		// If there was a repeated pattern, write the encoded pattern to the output.
		if count > 0 {
			if count >= 2 {
				fmt.Fprintf(&output, "[%d %s]", count+1, string(currentPattern))
			} else {
				fmt.Fprint(&output, strings.Repeat(string(currentPattern), count+1))
			}

			count = 0

			if len(currentPattern) == 2 {
				i++
			}
			currentPattern = ""
		} else {
			// Write the character to the output.
			fmt.Fprint(&output, string(text[i]))
		}
	}

	return output.String()
}
