package dynamicprogramming

// 动态规划
func longestPalindromeSubseq(s string) int {
	if len(s) < 2 {
		return len(s)
	}
	dp := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]int, len(s))
		dp[i][i] = 1
	}
	for l := 1; l < len(s); l++ {
		for i := 0; i+l < len(s); i++ {
			j := i + l
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][len(s)-1]
}
