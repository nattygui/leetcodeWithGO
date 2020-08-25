package rotatestring

import "strings"

func rotateString(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}
	if strings.Contains(A+A, B) {
		return true
	}
	return false
}
