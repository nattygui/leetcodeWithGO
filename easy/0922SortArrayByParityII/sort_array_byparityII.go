package sortarray

func sortArrayByParityII(A []int) []int {
	length := len(A)
	even := 0
	odd := 1

	for even < length && odd < length {
		if A[even]%2 == 0 {
			even += 2
			continue
		}
		if A[odd]%2 == 1 {
			odd += 2
			continue
		}
		A[even], A[odd] = A[odd], A[even]
		even += 2
		odd += 2
	}
	return A
}
