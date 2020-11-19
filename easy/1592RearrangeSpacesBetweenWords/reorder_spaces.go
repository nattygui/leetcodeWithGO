package reorderspaces

import (
	"bytes"
)

func reorderSpaces(text string) string {
	words := make([]string, 0)
	spaces := 0
	for i := 0; i < len(text); i++ {
		if text[i] == ' ' {
			spaces++
			continue
		}
		j := i
		for text[j] != ' ' && j < len(text) {
			j++
		}
		words = append(words, text[i:j])
		i = j - 1
	}
	length := len(words)
	var results bytes.Buffer
	for _, word := range words {
		results.WriteString(word)
		for i := 0; i < spaces/(length-1); i++ {
			results.WriteByte(' ')
		}
	}
	for i := 0; i < spaces%(length-1); i++ {
		results.WriteByte(' ')
	}
	return results.String()
}
