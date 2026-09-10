package main
import (
	//"fmt"
	"strings"
)
func cleanInput(text string) []string  {
	var resultao []string
	matchIndex := strings.Index(text, " ")
	for matchIndex > -1 {
		//fmt.Printf("    -Appending: %v\n", strings.ToLower(text[:matchIndex]))
		if text[:matchIndex] != "" {
			resultao = append(resultao, strings.ToLower(text[:matchIndex]))
		}
		text = text[matchIndex + 1:]
		matchIndex = strings.Index(text, " ")
	}
	if text != "" {
		resultao = append(resultao, strings.ToLower(text))
	}
	return resultao
}