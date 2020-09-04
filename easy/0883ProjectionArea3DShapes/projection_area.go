package projectionarea

func projectionArea(grid [][]int) int {
	areaZ := 0
	areaX := 0
	areaY := 0

	maxY := make([]int, 50)
	for _, g := range grid {
		areaZ += len(g)
		maxX := 0
		for i, t := range g {
			if maxX < t {
				maxX = t
			}

			if maxY[i] < t {
				maxY[i] = t
			}
		}
		areaX += maxX
	}
	areaY = sum(maxY)
	return areaX + areaY + areaZ
}

func sum(nums []int) int {
	s := 0
	for _, num := range nums {
		s += num
	}
	return s
}
