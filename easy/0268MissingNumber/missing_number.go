package _268MissingNumber

func missingNumber(nums []int) int {
	sum := 0
	for i := 1; i <= len(nums); i++ {
		sum += i
		sum -= nums[i-1]
	}

	return sum
}
