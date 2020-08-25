package greatestletter

import "testing"

func TestOK(t *testing.T) {
	letters := []byte{'c', 'f', 'j'}
	target := byte('j')

	res := byte('c')

	answer := nextGreatestLetter(letters, target)

	if res != answer {
		t.Error("error")
	}
}
