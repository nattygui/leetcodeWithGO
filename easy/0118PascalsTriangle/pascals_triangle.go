package _118PascalsTriangle

func generate(numRows int) [][]int {
	if numRows == 0 {
		return nil
	}
	if numRows == 1 {
		return [][]int{{1}}
	}
	if numRows == 2 {
		return [][]int{{1}, {1, 1}}
	}
	res := make([][]int, numRows)
	res[0] = []int{1}
	res[1] = []int{1, 1}
	for i := 2; i < numRows; i++ {
		temp := make([]int, i+1)
		temp[0] = 1
		temp[i] = 1
		for j := 1; j < i; j++ {
			temp[j] = res[i-1][j-1] + res[i-1][j]
		}
		res[i] = temp
	}
	return res
}
