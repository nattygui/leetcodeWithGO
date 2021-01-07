package medium

import (
	"strings"
)

func removeDuplicateLetters(s string) string {
	count := make([]int, 26)
	inStack := make([]bool, 26)
	stack := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		c := s[i]
		count[c-'a']++
	}

	for i := 0; i < len(s); i++ {
		c := s[i]
		index := c - 'a'
		count[index]--
		if inStack[index] {
			continue
		}

		lastIndex := len(stack) - 1
		for {
			if lastIndex == -1 {
				break
			}
			if count[stack[lastIndex]-'a'] == 0 {
				break
			}
			if stack[lastIndex] < c {
				break
			}
			if stack[lastIndex] >= c {
				inStack[stack[lastIndex]-'a'] = false
				stack = append(stack[:lastIndex], stack[lastIndex+1:]...)
				lastIndex--
			}
		}

		stack = append(stack, c)
		inStack[index] = true
	}

	var bs strings.Builder
	bs.Write(stack)
	return bs.String()
}
