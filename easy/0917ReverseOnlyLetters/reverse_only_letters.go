package reverseonlyletters

func reverseOnlyLetters(S string) string {
	symbols := []byte(S)

	start := 0
	end := len(S) - 1
	for start < end {
		if !isLetter(symbols[start]) {
			start++
		} else if !isLetter(symbols[end]) {
			end--
		} else {
			symbols[start], symbols[end] = symbols[end], symbols[start]
			start++
			end--
		}
	}

	return string(symbols)
}

func isLetter(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z')
}
