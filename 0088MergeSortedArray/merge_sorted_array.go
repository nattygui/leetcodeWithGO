package _088MergeSortedArray

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	length := len(nums1)
	for n > 0 {
		if m == 0 || nums1[m-1] <= nums2[n-1] {
			length--
			n--
			nums1[length] = nums2[n]
		} else {
			length--
			m--
			nums1[length] = nums1[m]
		}
	}
	return nums1
}
