package _350IntersectionOfTwoArraysII

import "sort"

// 使用map
func intersect(nums1 []int, nums2 []int) []int {
	maps := make(map[int]int, 0)
	res := make([]int ,0)

	for _, v := range nums1 {
		maps[v]++
	}

	for _, v := range nums2 {
		if maps[v] > 0 {
			res = append(res, v)
			maps[v]--
		}
	}
	return res
}

// 使用排序预处理
func intersectOne(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	res := make([]int, 0)
	i := 0
	j := 0

	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			i++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			res = append(res, nums1[i])
		}
	}
	return res
}