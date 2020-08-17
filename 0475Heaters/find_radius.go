package _475Heaters

import (
	"math"
	"sort"
)

func findRadius(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)

	housesStartIndex := 0
	maxRadius := 0
	i := 0
	j := 0




	for i < len(houses) && j < len(heaters) {
		if heaters[j] >= houses[i-1]{
			if j != 0  {
				temp := getMinRadius(houses[housesStartIndex: i], heaters[j], heaters[j-1])
				//if (heaters[j] - heaters[j - 1]) / 2 < heaters[j] - houses[housesStartIndex] {
				//	if maxRadius < (heaters[j] - heaters[j - 1]) / 2 {
				//		maxRadius = (heaters[j] - heaters[j - 1]) / 2
				//	}
				//} else {
				//	if maxRadius < heaters[j] - houses[housesStartIndex] {
				//		maxRadius = heaters[j] - houses[housesStartIndex]
				//	}
				//}
				if maxRadius < temp {
					maxRadius = temp
				}
			} else {
				if maxRadius < heaters[j] - houses[housesStartIndex]{
					maxRadius = heaters[j] - houses[housesStartIndex]
				}
			}
			i++
			j++
			housesStartIndex = i
		} else if heaters[j] < houses[i-1] {
			j++
		} else {
			i++
		}

	}

	if j == len(heaters) && i != len(houses) {
		temp := houses[len(houses)-1] - heaters[j-1]
		if temp > maxRadius {
			maxRadius = temp
		}
	}

	if i == len(houses) && j != len(heaters) {
		if j != 0 {
			temp := getMin3(heaters[j]-houses[housesStartIndex], houses[i-1]-heaters[j-1], (heaters[j]-heaters[j-1]) / 2)
			if temp > maxRadius {
				maxRadius = temp
			}
		} else {
			temp := heaters[j] - houses[housesStartIndex]
			if temp > maxRadius {
				maxRadius = temp
			}
		}
	}

	return maxRadius
}

func getMin2(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func getMin3(x, y, z int) int {
	return getMin2(getMin2(x,y), z)
}

func getMinRadius(houses []int, rend, rstart int) int {
	l := len(houses)
	left := 0
	right := 0
	if l % 2 == 1 {
		left = houses[l/2]
		right = houses[l/2]
	} else {
		left = houses[l/2-1]
		right = houses[l/2]
	}

	r1 := (rend - rstart) / 2
	r2 := rend - houses[0]
	r3 := houses[l-1] -rstart
	r4 := left - rstart
	r5 := rend - right
	return getMin3(getMin3(r1, r2, r3), r4, r5)
}


func findRadiusOne(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)
	low, high := 0, 1000000000
	for low < high {
		mid := low + (high - low) / 2
		hIdx := 0
		left, right := heaters[hIdx] - mid, heaters[hIdx] + mid
		var i int
		for i = 0; i < len(houses); {
			if houses[i] < left {
				break
			} else if houses[i] > right {
				hIdx++
				if hIdx >= len(heaters) {
					break
				}
				left, right = heaters[hIdx] - mid, heaters[hIdx] + mid
				continue
			}
			i++
		}
		if i == len(houses) {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}

func findRadiusTwo(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)

	heaters = append([]int{math.MinInt64/2}, heaters...)
	heaters = append(heaters, math.MaxInt64/2)

	maxRadius := 0
	tag := 0

	for i := 0; i < len(heaters); i++ {
		for j := tag; j < len(houses); j++ {
			if houses[j] <= heaters[i] {
				temp := getMinRadius(houses[tag:j+1], heaters[i], heaters[i-1])
				if maxRadius < temp {
					maxRadius = temp
				}
				tag = j + 1
			}
		}

	}
	return maxRadius
}

func findRadiusThere(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)

	heaters = append([]int{math.MinInt64/2}, heaters...)
	heaters = append(heaters, math.MaxInt64/2)

	maxRadius := 0
	i := 0

	for _, v := range houses {
		for i < len(heaters) - 1 && v > heaters[i] {
			i++
		}
		maxRadius = getMax(maxRadius, getMin(v-heaters[i-1], heaters[i]-v))
	}

	return maxRadius
}

func getMin(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func getMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}