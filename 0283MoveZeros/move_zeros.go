package _283MoveZeros

func moveZeroes(nums []int)  {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
	}

	for j < len(nums) {
		nums[j] = 0
	}
}
