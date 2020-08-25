package uniquemorserepresentations

func uniqueMorseRepresentations(words []string) int {
	morses := []string{
		".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--..",
	}
	resMap := make(map[string]int)
	for _, word := range words {
		morseStr := tranformMorse(word, morses)
		if _, ok := resMap[morseStr]; ok {
			resMap[morseStr]++
		} else {
			resMap[morseStr] = 1
		}
	}

	return len(resMap)
}

func tranformMorse(word string, morses []string) string {
	res := ""
	for _, letter := range word {
		res += morses[int(letter-'a')]
	}
	return res
}
