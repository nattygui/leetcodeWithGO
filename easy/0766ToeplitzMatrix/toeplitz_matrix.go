package toeplitzmatrix

func isToeplitzMatrix(matrix [][]int) bool {
	sizeY := len(matrix)
	sizeX := len(matrix[0])
	for newx := 1; newx < sizeX; newx++ {
		value := matrix[sizeY-1][newx]
		x := newx
		y := sizeY - 1
		for x >= 0 && y >= 0 {
			if value != matrix[y][x] {
				return false
			}
			x--
			y--
		}
	}

	for newy := 1; newy < sizeY; newy++ {
		value := matrix[newy][sizeX-1]
		x := sizeX - 1
		y := newy
		for x >= 0 && y >= 0 {
			if value != matrix[y][x] {
				return false
			}
			x--
			y--
		}
	}

	return true
}
