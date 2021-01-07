package kthlargestrelement

import (
	"math/rand"
	"time"
)

func findKthLargest(nums []int, k int) int {
	rand.Seed(time.Now().UnixNano())
	length := len(nums)
	left, right := 0, len(nums)-1
	target := length - k

	for {
		index := partition(nums, left, right)
		if index == target {
			return nums[index]
		} else if index < target {
			left = index + 1
		} else {
			right = index - 1
		}
	}
}

func partition(nums []int, left, right int) int {
	// 在区间随机选择一个元素作为标记点
	if right > left {
		randomIndex := left + 1 + rand.Intn(right-left)
		nums[left], nums[randomIndex] = nums[randomIndex], nums[left]
	}

	pivot := nums[left]

	lt, rt := left+1, right
	for {
		for lt <= rt && nums[lt] < pivot {
			lt++
		}
		for lt <= rt && nums[rt] > pivot {
			rt--
		}
		if lt > rt {
			break
		}
		nums[lt], nums[rt] = nums[rt], nums[lt]
		lt++
		rt--
	}
	nums[left], nums[rt] = nums[rt], nums[left]
	return rt
}
