package numsquares

func numSquares(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i < n+1; i++ {
		dp[i] = i
		for j := 1; i-j*j >= 0; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]
}

func min(num1, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}
