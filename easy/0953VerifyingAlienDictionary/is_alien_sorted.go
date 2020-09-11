package aliensorted

func isAlienSorted(words []string, order string) bool {
	letterMaps := make(map[byte]int, 26)
	for index := 0; index < 26; index++ {
		letterMaps[order[index]] = index
	}

	for index := 1; index < len(words); index++ {
		preWord := words[index-1]
		curWord := words[index]
		if !isValidSorted(preWord, curWord, letterMaps) {
			return false
		}
	}
	return true
}

func isValidSorted(preWord, curWord string, letterMaps map[byte]int) bool {
	preLens := len(preWord)
	curLens := len(curWord)

	i := 0

	for i < preLens && i < curLens && letterMaps[preWord[i]] == letterMaps[curWord[i]] {
		i++
	}
	if i == preLens {
		return true
	}

	if i == curLens {
		return false
	}

	if letterMaps[preWord[i]] < letterMaps[curWord[i]] {
		return true
	}

	return false
}
