package _409LongestPalindrome

func longestPalindrome(s string) int {
	countLetter := make([]int, 52)

	for _, v := range s {
		if v - 'A' >= 0 {
			countLetter[26 + v - 'A']++
		} else {
			countLetter[v - 'a']++
		}
	}

	sum := 0
	palindrome := 0

	for _, v := range countLetter {
		sum += v
		if v > 1 {
			palindrome += v / 2
		}
	}

	if sum > palindrome * 2 {
		return palindrome + 1
	}
	return palindrome
}
