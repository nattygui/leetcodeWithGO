package countnegatives

func countNegatives(grid [][]int) int {
	result := 0
	for i := 0; i < len(grid); i++ {
		length := len(grid[i])
		for j := 0; j < length; j++ {
			if grid[i][j] < 0 {
				result += length - j
				break
			}
		}
	}
	return result
}
