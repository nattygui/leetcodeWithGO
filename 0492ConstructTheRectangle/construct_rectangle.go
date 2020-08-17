package _492ConstructTheRectangle

import "math"

func constructRectangle(area int) []int {
	if area <= 0 {
		return nil
	}

	if area == 1 {
		return []int{1, 1}
	}

	min, max := 1, 0
	res := make([]int, 2)
	devition := math.MaxInt64

	for min <= area / 2 {
		if area % min == 0 {
			max = area / min
			if max >= min && (max - min) < devition  {
				devition = max - min
				res[0] = max
				res[1] = min
			}
		}
		min++
	}
	return res
}
