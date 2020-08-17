package _119PascalsTriangleII

// 直接建立一个二维切片，用来存储每一层的所有值，返回最后一层
func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	if rowIndex == 1 {
		return []int{1, 1}
	}
	rows := make([][]int, rowIndex+1)
	rows[0] = []int{1}
	rows[1] = []int{1, 1}
	for i := 2; i <= rowIndex; i++ {
		temp := make([]int, i+1)
		temp[0] = 1
		temp[i] = 1
		for j := 1; j < i; j++ {
			temp[j] = rows[i-1][j-1] + rows[i-1][j]
		}
		rows[i] = temp
	}
	return rows[rowIndex]
}

// 使用动态规划来计算
func getRowTwo(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	if rowIndex == 1 {
		return []int{1, 1}
	}

	row := make([]int, rowIndex+1)
	for i := 0; i < len(row); i++ {
		row[i] = 1
	}
	for i := 2; i <= rowIndex; i++ {
		for j := i-1; j > 0; j-- {
			row[j] = row[j-1] + row[j]
		}
	}
	return row
}