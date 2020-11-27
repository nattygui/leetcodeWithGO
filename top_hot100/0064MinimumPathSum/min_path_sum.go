package minpathsum

func minPathSum(grid [][]int) int {
	// 初始化
	rows := len(grid)
	cols := len(grid[0])
	for i := 1; i < rows; i++ {
		grid[i][0] += grid[i-1][0]
	}
	for j := 1; j < rows; j++ {
		grid[0][j] += grid[0][j-1]
	}
	// 线性加
	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {
			grid[i][j] += max(grid[i][j-1], grid[i-1][j])
		}
	}

	return grid[rows-1][cols-1]
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
