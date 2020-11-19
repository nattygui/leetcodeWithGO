package finalprices

func finalPrices(prices []int) []int {
	length := len(prices)
	for i := 0; i < length; i++ {
		nearMin := findNearMin(prices[i+1:], prices[i])
		if nearMin != -1 {
			prices[i] -= nearMin
		}
	}
	return prices
}

func findNearMin(prices []int, curPrice int) int {
	result := -1
	for _, price := range prices {
		if price <= curPrice {
			result = price
			break
		}
	}
	return result
}
