package _628MaximumProductofThreeNumbers

import "sort"

func maximumProduct(nums []int) int {
	sort.Ints(nums)

	num1 := nums[0] * nums[1] * nums[len(nums)-1]
	num2 := nums[len(nums)-1] * nums[len(nums)-2] * nums[len(nums)-3]

	if num1 > num2 {
		return num1
	}
	return num2
}
