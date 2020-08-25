package shortestcompletingword

func shortestCompletingWord(licensePlate string, words []string) string {
	str := make([]int, 26)
	for i := 0; i < len(licensePlate); i++ {
		key := licensePlate[i]
		switch {
		case key >= 'a' && key <= 'z':
			num := int(key - 'a')
			str[num]++
		case key >= 'A' && key <= 'Z':
			num := int(key - 'A')
			str[num]++
		}
	}

	results := []string{}

	for _, word := range words {
		tempStr := []int{}
		tempStr = append(tempStr, str...)
		if containAllLetter(word, tempStr) {
			results = append(results, word)
		}
	}
	if len(results) == 0 {
		return ""
	}

	minLen := 16
	minIndex := -1

	for i, result := range results {
		if len(result) < minLen {
			minLen = len(result)
			minIndex = i
		}
	}

	return results[minIndex]
}

func containAllLetter(word string, str []int) bool {
	for i := 0; i < len(word); i++ {
		str[int(word[i]-'a')]--
	}
	for _, v := range str {
		if v > 0 {
			return false
		}
	}
	return true
}
