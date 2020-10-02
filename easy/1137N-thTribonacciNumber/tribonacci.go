package tribonacci

func tribonacci(n int) int {
	if n == 1 {
		return 0
	}
	if n == 2 || n == 3 {
		return 1
	}
	nums := make([]int, n)
	copy(nums, []int{0, 1, 1})

	for i := 3; i < n; i++ {
		nums[i] = nums[i-1] + nums[i-2] + nums[i-3]
	}

	return nums[n-1]
}
