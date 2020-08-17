package longestworddictionary

import "testing"

func TestOK(t *testing.T) {
	words := []string{"m", "mo", "moc", "moch", "mocha", "l", "la", "lat", "latt", "latte", "c", "ca", "cat"}

	words = []string{"a", "banana", "app", "appl", "ap", "apply", "apple"}
	res := longestWord(words)
	if res != "apple" {
		t.Error("result error")
	}
}
