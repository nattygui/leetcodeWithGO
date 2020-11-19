package specialarray

import "sort"

func specialArray(nums []int) int {
	sort.Ints(nums)

	length := len(nums)
	for result := 1; result <= length; result++ {
		if nums[length-result] < result {
			continue
		}
		if length-result == 0 {
			return result
		}
		if nums[length-result-1] < result {
			return result
		}
	}

	return -1
}
