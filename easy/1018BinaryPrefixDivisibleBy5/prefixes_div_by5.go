package prefixesdivby5

func prefixesDivBy5(A []int) []bool {
	result := make([]bool, len(A))
	sum := 0
	for index, num := range A {
		sum = sum*2 + num
		sum %= 5
		if sum == 0 {
			result[index] = true
		}
	}
	return result
}
