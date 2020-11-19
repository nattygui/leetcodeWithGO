package isprefixofword

import "strings"

func isPrefixOfWord(sentence string, searchWord string) int {
	result := -1
	words := strings.Split(sentence, " ")
	for i, word := range words {
		if strings.HasPrefix(word, searchWord) {
			result = i + 1
			break
		}
	}
	return result
}
