package relativesortarray

import "sort"

type sortInt struct {
	arr   []int
	score map[int]int
}

func (si *sortInt) Len() int      { return len(si.arr) }
func (si *sortInt) Swap(i, j int) { si.arr[i], si.arr[j] = si.arr[j], si.arr[i] }
func (si *sortInt) Less(i, j int) bool {
	v1, v2 := si.score[si.arr[i]], si.score[si.arr[j]]
	if v1 == 0 && v2 == 0 {
		return si.arr[i] < si.arr[j]
	}
	return v1 < v2
}

func relativeSortArray(arr1 []int, arr2 []int) []int {
	score := make(map[int]int)
	for i, v := range arr2 {
		score[v] = i
	}
	si := &sortInt{
		arr:   arr1,
		score: score,
	}
	sort.Sort(si)

	return si.arr
}
