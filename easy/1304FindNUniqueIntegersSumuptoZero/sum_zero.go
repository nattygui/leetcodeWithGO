package sumzero

func sumZero(n int) []int {
	result := make([]int, n)
	num := n / 2
	for i := 0; i < num; i++ {
		result[2*i] = i + 1
		result[2*i+1] = -(i + 1)
	}
	if n%2 != 0 {
		result[n-1] = 0
	}
	return result
}
