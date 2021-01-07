package dynamicprogramming

func rob2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	dp1 := make([]int, len(nums)-1)
	dp1[0] = nums[0]
	dp1[1] = max(nums[0], nums[1])
	dp2 := make([]int, len(nums)-1)
	dp2[0] = nums[1]
	dp2[1] = max(nums[1], nums[2])
	newNums := nums[1:]
	for i := 2; i < len(nums)-1; i++ {
		dp1[i] = max(dp1[i-1], dp1[i-2]+nums[i])
		dp2[i] = max(dp2[i-1], dp2[i-2]+newNums[i])
	}
	length := len(dp1)
	return max(max(dp1[length-1], dp1[length-2]), max(dp2[length-1], dp2[length-2]))
}
