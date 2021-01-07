package numislands

var directs [4][2]int = [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func numIslands(grid [][]byte) int {
	result := 0
	for x, row := range grid {
		for y, num := range row {
			if num == '2' {
				continue
			}
			if num == '1' {
				grid[x][y] = '2'
				findAll(x, y, grid)
				result++
			}
		}
	}
	return result
}

func findAll(x, y int, grid [][]byte) {
	for _, direct := range directs {
		newx := x + direct[0]
		if newx < 0 || newx >= len(grid) {
			continue
		}
		newy := y + direct[1]
		if newy < 0 || newy >= len(grid[0]) {
			continue
		}
		if grid[newx][newy] == '1' {
			grid[newx][newy] = '2'
			findAll(newx, newy, grid)
		}
	}
}
