package mintimetovisitallpoints

func minTimeToVisitAllPoints(points [][]int) int {
	if len(points) <= 1 {
		return 0
	}
	result := 0
	xdirect := 1
	ydirect := 1
	for i := 1; i < len(points); i++ {
		points1 := points[i]
		points0 := points[i-1]
		if points1[0] > points0[0] {
			xdirect = 1
		} else {
			xdirect = -1
		}
		if points1[1] > points0[1] {
			ydirect = 1
		} else {
			ydirect = -1
		}
		for points1[1] != points0[1] && points1[0] != points0[0] {
			points0[0] += xdirect
			points0[1] += ydirect
			result++
		}
		if points1[0] != points0[0] {
			result += xdirect * (points1[0] - points0[0])
		}
		if points1[1] != points0[1] {
			result += ydirect * (points1[1] - points0[1])
		}
	}
	return result
}
