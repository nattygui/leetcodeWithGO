package minsubsequence

import "sort"

func minSubsequence(nums []int) []int {
	sort.Ints(nums)
	sums := sum(nums)
	result := make([]int, 0)
	tempSum := 0
	length := len(nums) - 1
	for {
		tempSum += nums[length]
		sums -= nums[length]
		result = append(result, nums[length])
		if tempSum > sums {
			break
		}
		length--
	}
	return result
}

func sum(nums []int) int {
	result := 0
	for _, num := range nums {
		result += num
	}
	return result
}
