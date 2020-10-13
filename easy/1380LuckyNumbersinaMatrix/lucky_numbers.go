package luckynumbers

func luckyNumbers(matrix [][]int) []int {
	result := make([]int, 0, len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		row := matrix[i]
		low := 100001
		col := -1
		for index, num := range row {
			if low > num {
				low = num
				col = index
			}
		}
		// 查找当前值是否是当前列最大的
		flag := true
		for j := 0; j < len(matrix); j++ {
			if low < matrix[j][col] {
				flag = false
			}
		}
		if flag {
			result = append(result, low)
		}
	}
	return result
}
