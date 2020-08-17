package longestworddictionary

func longestWord(words []string) string {
	res := ""
	stringsMap := make(map[string]bool)
	for _, word := range words {
		stringsMap[word] = true
	}

	flag := true
	for _, word := range words {
		tempWord := ""
		for i := 1; i < len(word); i++ {
			tempWord = word[0:i]
			if stringsMap[tempWord] != true {
				flag = false
			}
		}
		if flag && len(word) >= len(res) {
			if len(word) > len(res) {
				res = word
			} else {
				res = calBytes(res, word)
			}
		}
		flag = true
	}
	return res
}

func calBytes(str1, str2 string) string {
	i := 0
	lens := len(str1)
	for i < lens {
		if str1[i] == str2[i] {
			i++
			continue
		} else {
			if str1[i] > str2[i] {
				return str2
			} else {
				return str1
			}
		}
	}
	return str1
}
