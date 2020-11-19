package threesum

import "sort"

func threeSumWithMap(nums []int) [][]int {
	sort.Ints(nums)
	// 构建numsMap
	numsMap := make(map[int]int)
	for index, num := range nums {
		numsMap[num] = index
	}
	// 开始遍历
	result := make([][]int, 0)
	for index := 0; index < len(nums)-2; index++ {
		if index > 0 && nums[index] == nums[index-1] {
			continue
		}

		for index1 := index + 1; index1 < len(nums)-1; index1++ {
			if index1 > index+1 && nums[index1] == nums[index1-1] {
				continue
			}
			if v, ok := numsMap[-nums[index]-nums[index1]]; ok && index1 < v {
				result = append(result, []int{nums[index], nums[index1], -nums[index] - nums[index1]})
			}
		}
	}
	return result
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	length := len(nums)
	// 开始遍历
	var start, end int
	result := make([][]int, 0)
	for index := 0; index < length-2; index++ {
		if index > 0 && nums[index] == nums[index-1] {
			continue
		}
		start = index + 1
		end = length - 1
		for start < end {
			if start > index+1 && nums[start] == nums[start-1] {
				start++
				continue
			}
			if end < length-2 && nums[end] == nums[end+1] {
				end--
				continue
			}
			if nums[index]+nums[start]+nums[end] > 0 {
				end--
			} else if nums[index]+nums[start]+nums[end] < 0 {
				start++
			} else {
				result = append(result, []int{nums[index], nums[start], nums[end]})
				end--
				start++
			}
		}
	}
	return result
}
