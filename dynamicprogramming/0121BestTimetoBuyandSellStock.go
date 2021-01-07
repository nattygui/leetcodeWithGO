package dynamicprogramming

// 查找最近的最优解
func maxProfit1(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	result := 0
	minPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		if minPrice < prices[i] {
			if result < prices[i]-minPrice {
				result = prices[i] - minPrice
			}
		} else {
			minPrice = prices[i]
		}
	}
	return result
}

// 动态规划
func maxProfit2(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	dp := make([][2]int, len(prices))
	dp[0][0] = 0          // 卖出
	dp[0][1] = -prices[0] // 买入
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], -prices[i])
	}
	return dp[len(prices)-1][0]
}
