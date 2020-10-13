package findnumbers

func findNumbers(nums []int) int {
	result := 0
	for _, num := range nums {
		if isOddNumber(num) {
			result++
		}
	}
	return result
}

func isOddNumber(num int) bool {
	result := 0
	for num != 0 {
		result++
		num /= 10
	}
	return result%2 == 0
}
