package floodfill

type coor struct {
	x, y int
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if newColor < 0 && newColor > 65535 {
		return image
	}
	sizeX := len(image) - 1
	sizeY := 2
	if sr < 0 || sr > sizeX || sc < 0 || sc > sizeY {
		return image
	}
	initValue := image[sr][sc]
	image[sr][sc] = newColor
	coorFlags := make(map[coor]bool)
	coorFlags[coor{sr, sc}] = true
	// 初始化坐标
	coordinates := make([][]int, 0)
	for i := -1; i < 2; i = i + 2 {
		tempSr := sr + i
		if tempSr < 0 || tempSr > sizeX {
			continue
		}
		if image[tempSr][sc] == initValue {
			coordinates = append(coordinates, []int{tempSr, sc})
			coorFlags[coor{tempSr, sc}] = true
		}
	}

	for j := -1; j < 2; j = j + 2 {
		tempSc := sc + j
		if tempSc < 0 || tempSc > sizeY {
			continue
		}
		if image[sr][tempSc] == initValue {
			coordinates = append(coordinates, []int{sr, tempSc})
			coorFlags[coor{sr, tempSc}] = true
		}
	}

	// 开始遍历
	for len(coordinates) != 0 {
		tempCoordinates := make([][]int, 0)
		for _, coordinate := range coordinates {
			sr, sc := coordinate[0], coordinate[1]
			if image[sr][sc] == initValue {
				image[sr][sc] = newColor
				for i := -1; i < 2; i = i + 2 {
					tempSr := sr + i
					if tempSr < 0 || tempSr > sizeX {
						continue
					}

					if !coorFlags[coor{tempSr, sc}] && image[tempSr][sc] == initValue {
						tempCoordinates = append(tempCoordinates, []int{tempSr, sc})
					}
				}
				for j := -1; j < 2; j = j + 2 {
					tempSc := sc + j
					if tempSc < 0 || tempSc > sizeY {
						continue
					}
					if !coorFlags[coor{sr, tempSc}] && image[sr][tempSc] == initValue {
						tempCoordinates = append(tempCoordinates, []int{sr, tempSc})
					}
				}
			}
		}
		coordinates = tempCoordinates
	}
	return image
}
