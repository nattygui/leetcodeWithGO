package dynamicprogramming

// 暴力解法
func countSubstrings(s string) int {
	result := 0
	for i := 0; i < len(s); i++ {
		result++
		for j := i + 1; j < len(s); j++ {
			if validatePalindrome(s, i, j) {
				result++
			}
		}
	}
	return result
}

// 动态规划
func countSubstrings1(s string) int {
	result := 0
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}

	for l := 0; l < len(s); l++ {
		for i := 0; i+l < len(s); i++ {
			j := i + l
			if l == 0 {
				dp[i][j] = true
			} else if l == 1 {
				if s[i] == s[j] {
					dp[i][j] = true
				}
			} else {
				if s[i] == s[j] && dp[i+1][j-1] {
					dp[i][j] = true
				}
			}
			if dp[i][j] {
				result++
			}
		}
	}
	return result
}
