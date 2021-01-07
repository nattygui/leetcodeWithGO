package lis

func lengthOfLIS(nums []int) int {
	result := 0
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		if result < dp[i] {
			result = dp[i]
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
