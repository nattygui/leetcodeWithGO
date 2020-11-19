package longestpalindrome

// 中心拓展法
func longestPalindrome(s string) string {
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		start1, end1 := findMostResult(s, i, i)
		start2, end2 := findMostResult(s, i, i+1)
		if end1-start1 > end-start {
			end = end1
			start = start1
		}
		if end2-start2 > end-start {
			end = end2
			start = start2
		}
	}
	return s[start : end+1]
}

func findMostResult(s string, i, j int) (int, int) {
	for ; i >= 0 && j < len(s) && s[i] == s[j]; i, j = i-1, j+1 {
	}
	return i + 1, j - 1
}

//暴力解法
// func longestPalindrome(s string) string {
// 	start, end := 0, 0
// 	max := 0
// 	for i := 0; i < len(s); i++ {
// 		for j := len(s) - 1; j >= 0; j-- {
// 			if isPalindrome(s[i:j+1]) && j+1-i > max {
// 				start = i
// 				end = j + 1
// 				max = end - start
// 			}
// 		}
// 	}
// 	return s[start:end]
// }

// func isPalindrome(s string) bool {
// 	i, j := 0, len(s)-1
// 	for i < j {
// 		if s[i] != s[j] {
// 			return false
// 		}
// 		i++
// 		j--
// 	}
// 	return true
// }
