package laststoneweight

import "sort"

func lastStoneWeight(stones []int) int {
	if len(stones) == 1 {
		return stones[0]
	}
	sort.Ints(stones)
	length := len(stones)
	for stones[length-2] != 0 {
		if stones[length-1] == stones[length-2] {
			stones[length-1] = 0
			stones[length-2] = 0
		} else {
			stones[length-1] -= stones[length-2]
			stones[length-2] = 0
		}
		sort.Ints(stones)
	}
	return stones[length-1]
}
