package rob

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	nums[2] += nums[0]
	length := len(nums)
	for i := 3; i < len(nums); i++ {
		nums[i] = max(nums[i-2]+nums[i], nums[i-3]+nums[i])
	}
	return max(nums[length-1], nums[length-2])
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
