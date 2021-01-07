package reconstruct

import "sort"

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] < people[j][0] {
			return true
		}
		if people[i][0] == people[j][0] {
			return people[i][1] > people[j][1]
		}
		return false

	})
	hasSorted := make([]bool, len(people))
	result := make([][]int, len(people))
	for _, p := range people {
		num := p[1] + 1
		for i, v := range hasSorted {
			if !v {
				num--
				if num == 0 {
					result[i] = p
					hasSorted[i] = true
				}
			}
		}
	}
	return result
}
