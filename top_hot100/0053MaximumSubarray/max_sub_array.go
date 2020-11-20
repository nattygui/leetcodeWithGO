package maxsubarray

func maxSubArray(nums []int) int {
	max := nums[0]
	prev := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			if prev+nums[i] < 0 {
				prev = nums[i]
			} else {
				prev = prev + nums[i]
			}
		} else {
			if prev > 0 {
				prev = nums[i] + prev
			} else {
				prev = nums[i]
			}
		}
		if max < prev {
			max = prev
		}
	}
	return max
}
