package _387FirstUniqueCharacterInAString

func firstUniqChar(s string) int {
	charList := make(map[byte]bool, 0)

	for i := 0; i < len(s); i++ {
		if _, ok := charList[s[i]]; ok {
			charList[s[i]] = false
		} else {
			charList[s[i]] = true
		}
	}

	for i := 0; i < len(s); i++ {
		if charList[s[i]] == true {
			return i
		}
	}

	return -1
}

func firstUniqCharOne(s string) int {
	charList := make(map[byte]int, 0)

	for i := 0; i < len(s); i++ {
		charList[s[i]]++
	}

	for i := 0; i < len(s); i++ {
		if charList[s[i]] == 1 {
			return i
		}
	}
	return -1
}
