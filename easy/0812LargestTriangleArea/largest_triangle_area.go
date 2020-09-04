package largesttrianglearea

import "math"

func largestTriangleArea(points [][]int) float64 {
	lens := len(points)
	ans := 0.0
	for i := 0; i < lens; i++ {
		for j := i + 1; j < lens; j++ {
			for k := j + 1; k < lens; k++ {
				ans = max(ans, getArea(points[i], points[j], points[k]))
			}
		}
	}
	return ans
}

func max(value1, value2 float64) float64 {
	if value1 > value2 {
		return value1
	}
	return value2
}

func getArea(p, q, r []int) float64 {
	return 0.5 * math.Abs(float64(p[0]*q[1]-q[0]*p[1]+q[0]*r[1]-r[0]*q[1]+r[0]*p[1]-p[0]*r[1]))
}
