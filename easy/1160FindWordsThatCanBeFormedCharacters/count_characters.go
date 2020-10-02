package countcharacters

func countCharacters(words []string, chars string) int {
	result := 0
	charMaps := make(map[byte]int)
	for i := 0; i < len(chars); i++ {
		charMaps[chars[i]]++
	}

	for _, word := range words {
		if hasGrasp(word, charMaps) {
			result += len(word)
		}
	}
	return result
}

func hasGrasp(word string, charMap map[byte]int) bool {
	wordMap := make(map[byte]int)
	for i := 0; i < len(word); i++ {
		wordMap[word[i]]++
	}

	for k, v := range wordMap {
		if v > charMap[k] {
			return false
		}
	}

	return true
}
