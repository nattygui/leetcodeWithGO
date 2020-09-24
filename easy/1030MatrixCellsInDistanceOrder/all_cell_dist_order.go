package allcelldistorder

import "sort"

type distanceSort struct {
	result [][]int
	r      int
	c      int
}

func (ds distanceSort) Len() int {
	return len(ds.result)
}

func (ds distanceSort) Less(i, j int) bool {
	r, c := ds.r, ds.c
	point1 := ds.result[i]
	point2 := ds.result[j]

	distance1 := getDistance(r, c, point1[0], point1[1])
	distance2 := getDistance(r, c, point2[0], point2[1])
	return distance1 < distance2
}

func (ds distanceSort) Swap(i, j int) {
	ds.result[i], ds.result[j] = ds.result[j], ds.result[i]
}

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	distances := make([][]int, 0)
	r := 0
	c := 0
	for r < R {
		for c < C {
			distances = append(distances, []int{r, c})
			c++
		}
		c = 0
		r++
	}
	ds := distanceSort{
		result: distances,
		r:      r0,
		c:      c0,
	}

	sort.Sort(ds)
	return ds.result
}

func getDistance(r int, c int, r1 int, c1 int) int {
	result := 0
	if r < r1 {
		result += r1 - r
	} else {
		result += r - r1
	}

	if c < c1 {
		result += c1 - c
	} else {
		result += c - c1
	}

	return result
}
