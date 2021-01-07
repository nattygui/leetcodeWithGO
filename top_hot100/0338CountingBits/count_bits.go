package countbits

func countBits(num int) []int {
	dp := make([]int, num+1)
	for i := 0; i <= num; i++ {
		if i == 0 {
			dp[i] = 0
		} else if i == 1 || i == 2 {
			dp[i] = 1
		} else {
			dp[i] = findValue(dp, i)
		}
	}
	return dp
}

func findValue(dp []int, i int) int {
	if i <= 2 {
		return dp[i]
	}
	if i/2 == 0 {
		return dp[i/2]
	}
	return dp[i-1] + 1
}
