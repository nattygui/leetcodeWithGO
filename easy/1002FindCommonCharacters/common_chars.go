package commonchars

func commonChars(A []string) []string {
	letterMap := make(map[rune]int)

	firstWord := A[0]
	for _, letter := range firstWord {
		letterMap[letter]++
	}

	for i := 1; i < len(A); i++ {
		tempLetterMap := make(map[rune]int)
		curWord := A[i]
		for _, letter := range curWord {
			if _, ok := letterMap[letter]; ok {
				tempLetterMap[letter]++
			}
		}

		// 求交集
		// 删除letterMap 存在  tempLetterMap 不存在的元素
		for key := range letterMap {
			if _, ok := tempLetterMap[key]; !ok {
				delete(letterMap, key)
			}
		}
		// 获取同时存在的最小值
		for key, value := range tempLetterMap {
			if v, ok := letterMap[key]; ok {
				letterMap[key] = mini(value, v)
			}
		}
	}

	res := make([]string, 0)
	for key, value := range letterMap {
		for value > 0 {
			res = append(res, string(key))
			value--
		}
	}
	return res
}

func mini(a, b int) int {
	if a < b {
		return a
	}
	return b
}
