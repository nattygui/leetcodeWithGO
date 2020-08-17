package _680ValidPalindromeII

func validPalindrome(s string) bool {
	res := true
	i, j := 0, len(s)-1
	for i < j {
		if s[i] == s[j] {
			i++
			j--
		} else {
			res = isPalindrome(s, i, j-1) || isPalindrome(s, i+1, j)
			break
		}
	}
	return res
}

func isPalindrome(s string, i, j int) bool {
	for i < j {
		if s[i] == s[j] {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}
