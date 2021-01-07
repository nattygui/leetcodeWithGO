package maxprofit

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	result := 0
	minPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		curPrice := prices[i]
		if minPrice > curPrice {
			minPrice = curPrice
			continue
		}
		if result < curPrice-minPrice {
			result = curPrice - minPrice
		}
	}
	return result
}
