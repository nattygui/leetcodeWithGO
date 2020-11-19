package minstartvalue

func minStartValue(nums []int) int {
	minSum := 1
	sum := 0
	for _, num := range nums {
		sum += num
		if minSum > sum {
			minSum = sum
		}
	}
	if minSum < 1 {
		return -minSum + 1
	}
	return 1
}
