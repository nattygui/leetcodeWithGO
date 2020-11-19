package maxproduct

func maxProduct(nums []int) int {
	firstMax := 0
	secondeMax := 0
	for _, num := range nums {
		if firstMax < num {
			firstMax, secondeMax = num, firstMax
		} else if secondeMax < num {
			secondeMax = num
		}
	}
	return (firstMax - 1) * (secondeMax - 1)
}
