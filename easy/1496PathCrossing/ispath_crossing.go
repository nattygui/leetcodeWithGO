package ispathcrossing

func isPathCrossing(path string) bool {
	allPath := [][]int{}
	allPath = append(allPath, []int{0, 0})
	x := 0
	y := 0
	for i := 0; i < len(path); i++ {
		diffx, diffy := getDirection(path[i])
		x += diffx
		y += diffy
		if hasExist(allPath, x, y) {
			return true
		}
		allPath = append(allPath, []int{x, y})
	}
	return false
}

func hasExist(allPath [][]int, x int, y int) bool {
	for _, path := range allPath {
		if path[0] == x && path[1] == y {
			return true
		}
	}
	return false
}

func getDirection(b byte) (int, int) {
	switch b {
	case 'N':
		return 0, 1
	case 'S':
		return 0, -1
	case 'E':
		return 1, 0
	case 'W':
		return -1, 0
	default:
		return 0, 0
	}
}
