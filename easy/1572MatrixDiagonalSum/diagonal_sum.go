package diagonalsum

func diagonalSum(mat [][]int) int {
	sum := 0
	rows := len(mat)
	for i, j := 0, rows-1; i < rows && j >= 0; i, j = i+1, j+1 {
		sum += mat[i][j] + mat[j][i]
	}
	if rows%2 == 1 {
		sum -= mat[rows/2][rows/2]
	}
	return sum
}
