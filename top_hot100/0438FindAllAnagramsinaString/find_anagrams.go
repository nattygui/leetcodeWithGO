package findanagrams

func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}
	pchars := construct(p)
	firstChars := construct(s[:len(p)])
	sign := make([]bool, len(s)-len(p)+1)
	result := make([]int, 0)
	for i := 0; len(s)-i >= len(p); i++ {
		if i == 0 {
			if sliceEqual(pchars, firstChars) {
				sign[i] = true
				result = append(result, i)
			}
			continue
		}

		firstChars[s[i-1]-'a']--
		firstChars[s[i+len(p)-1]-'a']++
		if sign[i-1] {
			if s[i-1] == s[i+len(p)-1] {
				sign[i] = true
				result = append(result, i)
			}
		} else {
			if sliceEqual(pchars, firstChars) {
				sign[i] = true
				result = append(result, i)
			}
		}
	}
	return result
}

func construct(str string) []int {
	chars := make([]int, 26)
	for i := 0; i < len(str); i++ {
		chars[str[i]-'a']++
	}
	return chars
}

func sliceEqual(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
