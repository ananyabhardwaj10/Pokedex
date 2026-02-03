package main 
import (
	"strings"
)

func cleanInput(text string) []string {
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)

	var words []string 
	words = strings.Fields(text)
	return words
}