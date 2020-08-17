package _053MaximumSubarray

import "math"

// 暴力解法
// 时间复杂度 O(nxn)
func maxSubArray0(nums []int) int {
	max := math.MinInt64

	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum > max {
				max = sum
			}
		}
	}
	return max
}

// 贪心算法
// 时间复杂度 O(n) 空间复杂度
func maxSubArray1(nums []int) int {
	res := nums[0]
	sum := 0
	for _, value := range nums {
		if sum > 0 {
			sum += value
		} else {
			sum = value
		}
		if res < sum {
			res = sum
		}
	}
	return res
}

// 动态规划
// 时间复杂度
func maxSubArray2(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	result := dp[0]
	for i := 1; i < len(nums); i++ {
		if (dp[i-1] + nums[i]) > nums[i] {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
		if result < dp[i] {
			result = dp[i]
		}
	}
	return result
}