package numtrees

func numTrees(n int) int {
	if n <= 0 {
		return 0
	}
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		all := 0
		for j := 1; j <= i; j++ {
			all += dp[j-1] * dp[i-j]
		}
		dp[i] = all
	}
	return dp[n]
}
