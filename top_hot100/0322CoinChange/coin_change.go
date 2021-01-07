package coinchange

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i-coin >= 0 && dp[i-coin] != amount+1 {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

func min(num1, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}
