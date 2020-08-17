package _349IntersectionOfTwoArrays

import "sort"

// 使用映射
func intersection(nums1 []int, nums2 []int) []int {
	maps := make(map[int]bool, 0)
	res := make([]int, 0)

	for _, v := range nums1 {
		maps[v] = true
	}

	for _, v := range nums2 {
		if maps[v] == true {
			res = append(res, v)
			maps[v] = false
		}
	}
	return res
}

// 先排序 在查找
func intersectionOne(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	saveValue := make(map[int]bool, 0)
	res := make([]int, 0)

	i := 0
	j := 0
	l1 := len(nums1)
	l2 := len(nums2)

	for i < l1 && j < l2 {
		if nums1[i] < nums2[j] {
			i++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			saveValue[nums1[i]] = true
			i++
			j++
		}
	}

	for key, _ := range saveValue {
		res = append(res, key)
	}
	return res
}
