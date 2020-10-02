package oddcells

func oddCells(n int, m int, indices [][]int) int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, m)
	}

	for _, index := range indices {
		for i := 0; i < m; i++ {
			matrix[index[0]][i]++
		}
		for j := 0; j < n; j++ {
			matrix[j][index[1]]++
		}
	}
	result := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j]%2 == 1 {
				result++
			}
		}
	}
	return result
}
