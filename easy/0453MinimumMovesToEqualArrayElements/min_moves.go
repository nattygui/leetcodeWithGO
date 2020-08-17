package _453MinimumMovesToEqualArrayElements

import "sort"

// 暴力解法
func minMoves(nums []int) int {
	isEquals := func(nums []int) bool {
		value := nums[0]
		for i := 1; i < len(nums); i++ {
			if nums[i] != value {
				return false
			}
		}
		return true
	}

	addValues := func(nums []int, noAddIndex int) {
		for i := 0; i < len(nums); i++ {
			if i != noAddIndex {
				nums[i]++
			}
		}
	}

	getMaxvalueIndex := func(nums []int) int {
		temp := nums[0]
		key := 0
		for i := 1; i < len(nums); i++ {
			if temp < nums[i] {
				temp = nums[i]
				key = i
			}
		}
		return key
	}

	count := 0

	for !(isEquals(nums)) {
		max := getMaxvalueIndex(nums)
		addValues(nums, max)
		count++
	}
	return count
}

// 动态规划
func minMovesOne(nums []int) int {
	sort.Ints(nums)
	moves := 0
	diff := 0
	for i := 1; i < len(nums); i++ {
		diff = moves + nums[i] - nums[i-1]
		nums[i] += moves
		moves += diff
	}

	return moves
}

// 数学公式
func minMovesTwo(nums []int) int {
	moves := 0
	min := nums[0]
	for i := 1; i < len(nums); i++ {
		if min > nums[i] {
			min = nums[i]
		}
	}

	for _, v := range nums {
		moves += v - min
	}

	return moves
}
