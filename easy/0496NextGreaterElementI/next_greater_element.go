package _496NextGreaterElementI

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	numsMap := make(map[int]int)

	for i := 0; i < len(nums2); i++ {
		for j := i+1; j < len(nums2); j++ {
			if nums2[j] > nums2[i] {
				numsMap[nums2[i]] = nums2[j]
				break
			}
		}
	}

	for i := 0; i < len(nums1); i++ {
		if v, ok := numsMap[nums1[i]]; ok {
			nums1[i] = v
		} else {
			nums1[i] = -1
		}
	}

	return nums1
}
