package _506RelativeRanks

import (
	"sort"
	"strconv"
)

func findRelativeRanks(nums []int) []string {
	temp := []int{}
	temp = append(temp, nums...)
	sort.Ints(nums)

	res := make([]string, len(nums))
	index := 0
	for j := len(nums)-1; j >= 0; j-- {
		for i := 0; i < len(temp); i++ {
			if temp[i] == nums[j] {
				switch index {
				case 0:
					res[i] = "Gold Medal"
				case 1:
					res[i] = "Silver Medal"
				case 2:
					res[i] = "Bronze Medal"
				default:
					res[i] = strconv.Itoa(index+1)
				}
				index++
			}
		}

	}

	return  res
}
