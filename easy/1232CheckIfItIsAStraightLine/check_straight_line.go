package checkstraightline

func checkStraightLine(coordinates [][]int) bool {
	if len(coordinates) <= 2 {
		return true
	}
	bechPointOne := coordinates[0]
	bechPointSlope := []int{coordinates[1][0] - bechPointOne[0], coordinates[1][1] - bechPointOne[1]}

	for i := 2; i < len(coordinates); i++ {
		diffy := coordinates[i][1] - bechPointOne[1]
		diffx := coordinates[i][0] - bechPointOne[0]

		if diffy*bechPointSlope[0] != diffx*bechPointSlope[1] {
			return false
		}
	}
	return true
}
