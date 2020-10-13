package sortstring

func sortString(s string) string {
	lettersNum := make([]int, 26)

	strLen := len(s)
	for i := 0; i < strLen; i++ {
		lettersNum[s[i]-'a']++
	}

	result := make([]byte, 0, strLen)
	for strLen != 0 {
		for i := 0; i < 26; i++ {
			if lettersNum[i] == 0 {
				continue
			}
			lettersNum[i]--
			result = append(result, byte(i+'a'))
			strLen--
		}

		for i := 25; i >= 0; i-- {
			if lettersNum[i] == 0 {
				continue
			}
			lettersNum[i]--
			result = append(result, byte(i+'a'))
			strLen--
		}
	}

	return string(result)
}
