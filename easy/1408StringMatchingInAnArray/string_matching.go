package stringmatching

import (
	"strings"
)

func stringMatching(words []string) []string {
	result := make([]string, 0)
	for i := 0; i < len(words); i++ {
		temp := words[i]
		for j := 0; j < len(words); j++ {
			if i != j && strings.Contains(words[j], temp) {
				result = append(result, temp)
				break
			}
		}
	}
	return result
}
