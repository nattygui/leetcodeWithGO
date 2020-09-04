package faircandyswap

func fairCandySwap(A []int, B []int) []int {
	sumA := sum(A)
	sumB := sum(B)

	res := []int{0, 0}
	diff := (sumB - sumA) / 2
	for _, x := range A {
		for _, y := range B {
			if y == x+diff {
				res[0] = x
				res[1] = y
				return res
			}
		}
	}
	return res
}

func sum(nums []int) int {
	s := 0
	for _, num := range nums {
		s += num
	}
	return s
}
