package maxproduct

func maxProduct(nums []int) int {
	result := nums[0]
	minDp := make([]int, len(nums))
	maxDp := make([]int, len(nums))
	for index, num := range nums {
		if index == 0 {
			minDp[0] = num
			maxDp[0] = num
			continue
		}
		minDp[index] = min(minDp[index-1]*num, min(maxDp[index-1]*num, num))
		maxDp[index] = max(maxDp[index-1]*num, max(minDp[index-1]*num, num))
		if result < maxDp[index] {
			result = maxDp[index]
		}
		if result < minDp[index] {
			result = minDp[index]
		}
	}
	return result
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}

func min(num1, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}
