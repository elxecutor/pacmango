package main

import (
	"os"
	"strconv"
	"strings"
)

const highScoreFile = "highscore.txt"

// LoadHighScore loads the high score from a file
func LoadHighScore() int {
	data, err := os.ReadFile(highScoreFile)
	if err != nil {
		return 0 // If file doesn't exist, start at 0
	}
	s := trimWhitespace(string(data))
	score, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return score
}

// trimWhitespace removes leading/trailing whitespace and newlines
func trimWhitespace(s string) string {
	return strings.TrimSpace(s)
}

// SaveHighScore saves the high score to a file
func SaveHighScore(score int) {
	f, err := os.OpenFile(highScoreFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return
	}
	defer f.Close()
	f.WriteString(strconv.Itoa(score))
}
