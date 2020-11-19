package frequencysort

import "sort"

type resultSort struct {
	nums     []int
	scoreMap map[int]int
}

func (rs *resultSort) Len() int {
	return len(rs.nums)
}
func (rs *resultSort) Less(i, j int) bool {
	if rs.scoreMap[rs.nums[i]] < rs.scoreMap[rs.nums[j]] {
		return true
	}
	if rs.scoreMap[rs.nums[i]] == rs.scoreMap[rs.nums[j]] && rs.nums[i] > rs.nums[j] {
		return true
	}
	return false
}
func (rs *resultSort) Swap(i, j int) {
	rs.nums[i], rs.nums[j] = rs.nums[j], rs.nums[i]
}

func frequencySort(nums []int) []int {
	scoreMap := make(map[int]int)
	for _, num := range nums {
		if _, ok := scoreMap[num]; ok {
			scoreMap[num]++
		} else {
			scoreMap[num] = 1
		}
	}
	resultScore := &resultSort{
		nums:     nums,
		scoreMap: scoreMap,
	}
	sort.Sort(resultScore)
	return resultScore.nums
}
