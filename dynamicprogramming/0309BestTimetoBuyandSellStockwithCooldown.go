package dynamicprogramming

func maxProfit4(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	dp := make([][2]int, len(prices))
	dp[0][0] = -prices[0] // 买入
	dp[0][1] = 0          // 卖出

	dp[1][0] = max(dp[0][0], dp[0][1]-prices[1])
	dp[1][1] = max(dp[0][1], dp[0][0]+prices[1])

	for i := 2; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-2][1]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}
	return dp[len(prices)-1][1]
}
