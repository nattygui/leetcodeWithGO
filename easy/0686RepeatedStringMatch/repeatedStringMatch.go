package _686RepeatedStringMatch

import "strings"

func repeatedStringMatch(A string, B string) int {
	la, lb := len(A), len(B)

	minlen := la*2 + lb
	t, res := "", 0
	for len(t) < minlen {
		if strings.Contains(t, B) {
			return res
		}
		res++
		t += A
	}
	return -1
}
