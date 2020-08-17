package _447NumberOfBoomerangs

func numberOfBoomerangs(points [][]int) int {
	count := 0

	computeDistance := func(x, y []int) int {
		return (x[0] - y[0]) * (x[0] - y[0]) + (x[1] - y[1]) * (x[1] - y[1])
	}

	for i := 0; i < len(points); i++ {
		point1 := points[i]
		records := make(map[int]int, 0)
		for j := i+1; j < len(points); j++ {
			point2 := points[j]
			if i != j {
				records[computeDistance(point1, point2)]++
			}
		}

		for _, v := range records {
			if v >= 2 {
				count += v * (v - 1)
			}
		}
	}

	return count
}


