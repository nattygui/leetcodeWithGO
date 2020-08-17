package _674LongestContinuousIncreasingSubsequence

func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}
	record := 1
	res := 0
	for i := 1; i < len(nums); i++ {
		temp := nums[i] - nums[i-1]
		if temp > 0 {
			record++
		} else {
			if res < record {
				res = record
			}
			record = 1
		}
	}
	if res < record {
		res = record
	}
	return res
}
