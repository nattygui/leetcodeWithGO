package searchmatrix

func searchMatrix(matrix [][]int, target int) bool {
	rows := len(matrix)
	cols := len(matrix[0])
	i := rows - 1
	j := 0
	for {
		if i < 0 || j > cols-1 {
			return false
		}
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] > target {
			i--
		} else if matrix[i][j] < target {
			j++
		}
	}
}
