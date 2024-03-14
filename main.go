package main

import (
	"bufio"
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
		action := getActionFromUser(reader)
		if action == "" {
			continue
		}

		inputType := getInputTypeFromUser(reader)
		if inputType == "" {
			continue
		}

		text := getTextArtFromUser(reader, inputType)
		if text == "" {
			continue
		}

		if filepath.Ext(strings.TrimSpace(text)) == ".txt" {
			text = handleTextFile(text)
		} else {
			//text = strings.TrimSpace(text)
		}

		if action == "e" {
			encodeText(text)
		} else if action == "d" {
			decodeText(text)
		}

		if !continueOperation(reader) {
			break
		}
	}
}
