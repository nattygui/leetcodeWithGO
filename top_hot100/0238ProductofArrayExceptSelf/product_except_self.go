package productexceptself

func productExceptSelf(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	if len(nums) == 1 {
		return []int{1}
	}
	answer := make([]int, len(nums))
	answer[0] = 1
	for i := 1; i < len(nums); i++ {
		answer[i] = answer[i-1] * nums[i-1]
	}

	r := 1
	for i := len(nums) - 2; i >= 0; i-- {
		r *= nums[i+1]
		answer[i] *= r
	}
	return answer
}
