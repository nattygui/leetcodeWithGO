package dynamicprogramming

// 暴力解法
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	return max(getMaxValue(nums, len(nums)-1), getMaxValue(nums, len(nums)-2))
}

func getMaxValue(nums []int, index int) int {
	if index == 0 {
		return nums[index]
	}
	if index == 1 {
		return max(nums[0], nums[1])
	}
	return max(nums[index]+getMaxValue(nums, index-2), getMaxValue(nums, index-1))
}

// 动态规划
func rob1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return max(dp[len(nums)-1], dp[len(nums)-2])
}
