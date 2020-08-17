package _122BestTimeToBuyAndSellStockII

// 暴力法，计算出所有可能的交易组合相对应的利润，并找出它们中的最大利润  时间复杂度O(n*2) 控件复杂度O(n)
func maxProfitOne(prices []int) int {
	return calculate(prices, 0)
}

func calculate(prices []int, s int) int {
	if s >= len(prices) {
		return 0
	}
	max := 0
	for start := s; start < len(prices); start++ {
		maxProfit := 0
		for i := start + 1; i < len(prices); i++ {
			if prices[start] < prices[i] {
				profit := calculate(prices, i+1) + prices[i] - prices[start]
				if profit>maxProfit{
					maxProfit = profit
				}
			}
		}
		if maxProfit > max {
			max = maxProfit
		}
	}
	return max
}

// 贪心算法，只要今天的价格大于昨天的。就卖出
func maxProfitTwo(prices []int) int {
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit = profit + prices[i] - prices[i-1]
		}
	}
	return profit
}

// 动态规划

func maxProfitThird(prices []int) int {
	// 初始化
	if len(prices) <= 1 {
		return 0
	}
	dp := make([][2]int, len(prices))
	dp[0][0] = 0
	dp[0][1] = -prices[0]

	// 状态转移
	for i := 1; i < len(prices); i++ {
		if dp[i-1][0] < dp[i-1][1] + prices[i] {
			dp[i][0] = dp[i-1][1] + prices[i]
		} else {
			dp[i][0] = dp[i-1][0]
		}

		if dp[i-1][1] <  dp[i-1][0] - prices[i] {
			dp[i][1] = dp[i-1][0] - prices[i]
		} else {
			dp[i][1] = dp[i-1][1]
		}
	}
	return dp[len(prices)][0]
}