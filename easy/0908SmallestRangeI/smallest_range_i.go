package smallestrangei

func smallestRangeI(A []int, K int) int {
	min, max := A[0], A[0]

	for _, x := range A {
		if min > x {
			min = x
		}
		if max < x {
			max = x
		}
	}

	res := (max - K) - (min + K)

	if res < 0 {
		return 0
	}
	return res
}
