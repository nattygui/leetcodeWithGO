package invertimage

func flipAndInvertImage(A [][]int) [][]int {
	for _, row := range A {
		invertImage(row)
	}
	return A
}

func invertImage(row []int) {
	i, j := 0, len(row)-1
	for i <= j {
		if i == j {
			row[i] = abs(row[i] - 1)
			return
		}
		row[i], row[j] = row[j], row[i]
		row[i] = abs(row[i] - 1)
		row[j] = abs(row[j] - 1)
		i++
		j--
	}
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
