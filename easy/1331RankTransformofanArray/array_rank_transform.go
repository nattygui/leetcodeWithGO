package arrayranktransform

import "sort"

func arrayRankTransform(arr []int) []int {
	values := make([]int, 0)

	indexMap := make(map[int][]int)
	for index, value := range arr {
		values = append(values, value)
		if indice, ok := indexMap[value]; ok {
			indice = append(indice, index)
		} else {
			indexMap[value] = []int{index}
		}
	}

	sort.Ints(values)
	for index, value := range values {
		for _, t := range indexMap[value] {
			arr[t] = index + 1
		}
	}
	return arr
}
