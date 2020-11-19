package numspecial

func numSpecial(mat [][]int) int {
	rows := len(mat)
	cols := len(mat[0])
	rowsMap := make(map[int]int)
	colsMap := make(map[int]int)
	onePoints := [][]int{}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if mat[i][j] == 1 {
				onePoints = append(onePoints, []int{i, j})
				rowsMap[i]++
				colsMap[j]++
			}
		}
	}
	results := 0
	for _, point := range onePoints {
		if rowsMap[point[0]] == 1 && colsMap[point[1]] == 1 {
			results++
		}
	}
	return results
}
