package wordbreak

func wordBreak(s string, wordDict []string) bool {
	numWordsMap := make(map[int][]string)
	// 构建数组map
	for _, word := range wordDict {
		length := len(word)
		if v, ok := numWordsMap[length]; !ok {
			numWordsMap[length] = []string{word}
		} else {
			numWordsMap[length] = append(v, word)
		}
	}

	return backtrack(s, numWordsMap)
}

func backtrack(s string, numWordsMap map[int][]string) bool {
	if len(s) == 0 {
		return true
	}
	result := false
	for i := 0; i < len(s); i++ {
		if words, ok := numWordsMap[i+1]; ok {
			word := s[:i+1]
			for _, temp := range words {
				if word == temp {
					result = backtrack(s[i+1:], numWordsMap) || result
				}
			}
		}
	}
	return result
}
