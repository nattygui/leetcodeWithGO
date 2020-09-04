package transpose

func transpose(A [][]int) [][]int {
	lenX := len(A[0])
	lenY := len(A)

	res := make([][]int, lenX)
	for i := 0; i < lenX; i++ {
		temp := make([]int, lenY)
		for j := 0; j < lenY; j++ {
			temp[j] = A[j][i]
		}
		res[i] = temp
	}

	return res
}
