package dynamicprogramming

// 使用暴力解法
func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	maxLen := 0
	start := 0
	end := 1
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if j-i+1 > maxLen && validatePalindrome(s, i, j) {
				maxLen = j - i + 1
				start = i
				end = j + 1
			}
		}
	}
	return s[start:end]
}

func validatePalindrome(s string, i, j int) bool {
	for i <= j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

// 动态规划
// 从暴力解法中得知，若判断s[i, j] 是不是回文，只需要判断 s[i-1, j-1] && s[i] == s[j]
// 所以可以构建二维矩阵 dp[n][n]
func longestPalindrome1(s string) string {
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	result := ""
	for l := 0; l < len(s); l++ {
		for i := 0; i+l < len(s); i++ {
			j := i + l
			if l == 0 {
				dp[i][j] = true
			} else if l == 1 {
				if s[i] == s[j] {
					dp[i][j] = true
				} else {
					dp[i][j] = false
				}
			} else {
				dp[i][j] = dp[i+1][j-1] && s[i] == s[j]
			}
			if dp[i][j] && j-i+1 > len(result) {
				result = s[i : j+1]
			}
		}
	}
	return result
}
