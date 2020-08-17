package _661ImageSmoother

func imageSmoother(M [][]int) [][]int {
	row, col := len(M), len(M[0])
	res := make([][]int, row)

	for i := 0; i < row; i++ {
		temp := make([]int, col)
		for j := 0; j < col; j++ {
			sum := 0
			count := 0
			for t1 := i-1; t1 <= i+1; t1++ {
				if t1 >= 0 && t1 <= row - 1 {
					for t2 := j-1; t2 <= j+1; t2++ {
						if t2 >= 0 && t2 <= col - 1 {
							count++
							sum += M[t1][t2]
						}
					}
				}
			}
			temp[j] = sum / count
		}
		res[i] = temp
	}
	return res
}
