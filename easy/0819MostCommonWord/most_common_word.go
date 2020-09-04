package mostcommonword

func mostCommonWord(paragraph string, banned []string) string {
	wordMap := make(map[string]int)

	word := []byte{}
	for i := 0; i < len(paragraph); i++ {
		b := paragraph[i]
		if b >= 'a' && b <= 'z' {
			word = append(word, b)
			continue
		} else if b >= 'A' && b <= 'Z' {
			word = append(word, b-'A'+'a')
			continue
		}

		if len(word) == 0 {
			continue
		}

		newStr := string(word)

		if strInStrings(newStr, banned) {
			word = []byte{}
			continue
		}

		if _, ok := wordMap[newStr]; ok {
			wordMap[newStr]++
		} else {
			wordMap[newStr] = 1
		}
		word = []byte{}
	}
	// 最后查找次数最多的单词
	max := 0
	res := ""
	for key, value := range wordMap {
		if max < value {
			res = key
			max = value
		}
	}
	return res
}

func strInStrings(str string, banned []string) bool {
	for _, v := range banned {
		if str == v {
			return true
		}
	}
	return false
}
