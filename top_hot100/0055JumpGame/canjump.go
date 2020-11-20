package canjump

func canJump(nums []int) bool {
	length := len(nums)
	maxRight := 0
	for i := range nums {
		if i > maxRight {
			return false
		}
		maxRight := max(maxRight, i+nums[i])
		if maxRight >= length-1 {
			return true
		}
	}
	return false
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
