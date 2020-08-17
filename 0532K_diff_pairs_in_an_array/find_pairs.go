package _532K_diff_pairs_in_an_array

import "sort"

func findPairs(nums []int, k int) int {
	sort.Ints(nums)
	res := 0

	if k != 0 {
		for i := 0; i < len(nums); i++ {
			if (i+1 < len(nums)) && nums[i] == nums[i+1] {
				continue
			}
			for j := i + 1; j < len(nums); j++ {
				if nums[j]-nums[i] == k {
					res++
					break
				}
			}
		}
	} else {
		for i := 0; i < len(nums); i++ {
			for j := i + 1; j < len(nums); j++ {
				if nums[j]-nums[i] == k {
					temp := nums[j]
					res++
					for j < len(nums) && nums[j] == temp {
						i++
						j++
					}
					break
				}
			}

		}
	}
	return res
}

// 使用map映射
func findPairs1(nums []int, k int) int {
	res := 0
	kvs := make(map[int]bool)

	if k == 0 {
		for _, v := range nums {
			if _, ok := kvs[v]; ok {
				kvs[v] = true
			} else {
				kvs[v] = false
			}
		}
		for _, v := range kvs {
			if v == true {
				res++
			}
		}
		return res
	}

	if k <= 0 {
		return 0
	}

	for _, v := range nums {
		kvs[v+k] = false
	}

	for _, v := range nums {
		if _, ok := kvs[v]; ok {
			kvs[v] = true
		}
	}

	for _, v := range kvs {
		if v == true {
			res++
		}
	}

	return res
}
