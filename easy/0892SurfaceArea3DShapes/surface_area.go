package surfacearea

func surfaceArea(grid [][]int) int {
	res := 0
	for i, g := range grid {
		for j, s := range g {
			if s == 0 {
				continue
			}
			res += s*4 + 2
			res -= getOverlap(i, j, grid)
		}
	}
	return res
}

func getOverlap(i, j int, grid [][]int) int {
	res := 0
	cur := grid[i][j]

	res += getValue(i-1, j, cur, grid)
	res += getValue(i+1, j, cur, grid)
	res += getValue(i, j-1, cur, grid)
	res += getValue(i, j+1, cur, grid)

	return res
}

func getValue(i, j, cur int, grid [][]int) int {
	if i < 0 || i >= len(grid) {
		return 0
	}

	lens := len(grid[i])

	if j < 0 || j >= lens {
		return 0
	}

	if cur < grid[i][j] {
		return cur
	}

	return grid[i][j]
}
