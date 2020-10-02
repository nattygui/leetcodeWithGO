package shiftgrid

func shiftGrid(grid [][]int, k int) [][]int {
	row := len(grid)
	col := len(grid[0])
	size := row * col
	k %= size

	if k == 0 {
		return grid
	}
	result := make([][]int, row)
	for i := 0; i < row; i++ {
		result[i] = make([]int, col)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			result[(i+(k+j)/col)%row][(j+k)%col] = grid[i][j]
		}
	}
	return result
}
