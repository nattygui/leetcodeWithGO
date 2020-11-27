package merge

import "sort"

type sortSliceInt [][]int

func (ss sortSliceInt) Len() int           { return len(ss) }
func (ss sortSliceInt) Less(i, j int) bool { return ss[i][0] < ss[j][0] }
func (ss sortSliceInt) Swap(i, j int)      { ss[i], ss[j] = ss[j], ss[i] }

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	sort.Sort(sortSliceInt(intervals))
	checkIndex := 0
	for i := 1; i < len(intervals); {
		if intervals[checkIndex][0] <= intervals[i][0] && intervals[checkIndex][1] >= intervals[i][0] {
			intervals[checkIndex][1] = max(intervals[checkIndex][1], intervals[i][1])
			intervals = append(intervals[:i], intervals[i+1:]...)
		} else {
			checkIndex++
			if checkIndex == i {
				i++
			}
		}
	}
	return intervals
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
