package findsolution

func findSolution(customFunction func(int, int) int, z int) [][]int {
	res := make([][]int, 0)
	i, j := 1, 1000
	for i <= 1000 && j >= 1 {
		if customFunction(i, j) == z {
			res = append(res, []int{i, j})
			i++
			j--
		} else if customFunction(i, j) < z {
			i++
		} else {
			j--
		}
	}
	return res
}
