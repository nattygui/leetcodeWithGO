package sortarray

func sortArrayByParity(A []int) []int {
	res := make([]int, len(A))

	evenIndex, oddIndex := 0, len(A)-1
	for _, s := range A {
		if s%2 == 0 {
			res[evenIndex] = s
			evenIndex++
		} else {
			res[oddIndex] = s
			oddIndex--
		}
	}
	return res
}

func sortArrayByParity1(A []int) []int {
	i := 0
	j := len(A) - 1

	for i < j {
		if A[i]%2 == 0 {
			i++
		} else if A[j]%2 == 1 {
			j--
		} else {
			A[i], A[j] = A[j], A[i]
			i++
			j--
		}
	}

	return A
}
