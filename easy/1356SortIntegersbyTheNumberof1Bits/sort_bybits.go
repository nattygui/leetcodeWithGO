package sortbybits

import "sort"

func sortByBits(arr []int) []int {
	countBits := make([]int, len(arr))
	for index, num := range arr {
		countBits[index] = countBit(num)
	}

	sortbits := &SortByBits{
		arr:       arr,
		bitsCount: countBits,
	}
	sort.Sort(sortbits)

	return sortbits.bitsCount
}

func countBit(num int) int {
	result := 0
	for num != 0 {
		if num%2 == 1 {
			result++
		}
		num /= 2
	}
	return result
}

// SortByBits 根据bit 位数的大小进行排序
type SortByBits struct {
	arr       []int
	bitsCount []int
}

func (a *SortByBits) Len() int { return len(a.arr) }
func (a *SortByBits) Swap(i, j int) {
	a.arr[i], a.arr[j] = a.arr[j], a.arr[i]
	a.bitsCount[i], a.bitsCount[j] = a.bitsCount[j], a.bitsCount[i]
}
func (a *SortByBits) Less(i, j int) bool {
	if a.bitsCount[i] == a.bitsCount[j] {
		return a.arr[i] < a.arr[j]
	}
	return a.bitsCount[i] < a.bitsCount[j]
}
