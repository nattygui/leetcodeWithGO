package _121BestTimeToBuyAndSellStock

// 直接做 时间复杂度为o(n*2) 空间复杂度o(n)
func maxProfit(prices []int) int {
	minStock := 0
	maxStock := 0
	maxProfitV := 0
	for i := 0; i < len(prices); i++ {
		minStock = prices[i]
		maxStock = prices[i]
		for j := i+1; j < len(prices); j++ {
			if maxStock < prices[j] {
				maxStock = prices[j]
			}
		}
		if maxProfitV < (maxStock - minStock) {
			maxProfitV = maxStock - minStock
		}
	}
	return maxProfitV
}

// 动态规划
/*
	基本思路方法
	1、 获取当日的利润 todayProfit
    2、 当日的 stockValue 与 当前的 minPrice 对比 获取最小的值
	3、 当前的最大 maxProfit 与 当日的利润进行对比 获取最大的 maxProfit
*/
func maxProfitTwo(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	minPriceValue := prices[0]
	maxProfitValue := 0
	todayProfit := 0

	for i := 0; i < len(prices); i++ {
		todayProfit = prices[i] - minPriceValue
		if maxProfitValue < todayProfit {
			maxProfitValue = todayProfit
		}
		if minPriceValue > prices[i] {
			minPriceValue = prices[i]
		}
	}

	return maxProfitValue
}