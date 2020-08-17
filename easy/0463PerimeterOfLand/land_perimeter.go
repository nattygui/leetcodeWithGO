package _463PerimeterOfLand

func islandPerimeter(grid [][]int) int {
	getCellPerimeter := func(grid [][]int, x, y int) int {
		count := 0
		if x == 0 || grid[x-1][y] == 0 {
			count++
		}
		if x == len(grid) - 1 || grid[x+1][y] == 0 {
			count++
		}
		if y == 0 || grid[x][y-1] == 0 {
			count++
		}
		if y == len(grid[0]) - 1 || grid[x][y+1] == 0 {
			count++
		}
		return count
	}

	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				res += getCellPerimeter(grid, i, j)
			}
		}
	}

	return res
}
