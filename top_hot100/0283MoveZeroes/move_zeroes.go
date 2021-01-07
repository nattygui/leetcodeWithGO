package movezeroes

func moveZeroes(nums []int) {
	start := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[start] = nums[i]
			start++
			continue
		}
	}
	for start < len(nums) {
		nums[start] = 0
		start++
	}
}
