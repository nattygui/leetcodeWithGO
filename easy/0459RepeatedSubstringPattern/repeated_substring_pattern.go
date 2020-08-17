package _459RepeatedSubstringPattern

func repeatedSubstringPattern(s string) bool {
	isSubstringPattern := func(s, pattern string) bool {
		l1 := len(pattern)
		l2 := len(s)
		if l2 % l1 != 0 {
			return false
		}
		i := 1
		for i * l1 <= l2 {
			if s[l1*(i-1): l1*i] == pattern {
				i++
			} else {
				return false
			}
		}
		return true
	}


	for j := 1; j <= len(s) / 2; j++ {
		if s[j] == s[0] {
			if isSubstringPattern(s, s[0:j]) {
				return true
			}
		}
	}

	return false

}
