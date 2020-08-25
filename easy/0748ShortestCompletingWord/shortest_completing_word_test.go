package shortestcompletingword

import "testing"

func TestOK(t *testing.T) {
	words := []string{"suggest", "letter", "of", "husband", "easy", "education", "drug", "prevent", "writer", "old"}
	licensePlate := "Ah71752"

	res := shortestCompletingWord(licensePlate, words)
	if res != "husband" {
		t.Errorf("error")
	}
}
