package findmediansortedarrays

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	length1, lenght2 := len(nums1), len(nums2)

	firstTarget, secondTarget := 0, 0
	if (length1+lenght2)%2 == 0 {
		firstTarget = (length1 + lenght2) / 2
		secondTarget = (length1+lenght2)/2 + 1
	} else {
		firstTarget = (length1+lenght2)/2 + 1
		secondTarget = firstTarget
	}

	firstValue, _ := findKthElement(nums1, nums2, firstTarget)
	if firstTarget == secondTarget {
		return float64(firstValue)
	}
	secondValue, _ := findKthElement(nums1, nums2, secondTarget)

	return (float64(firstValue+secondValue) / 2.0)
}

func findKthElement(nums1 []int, nums2 []int, index int) (int, bool) {
	if len(nums1) == 0 {
		return nums2[index-1], true
	}
	if len(nums2) == 0 {
		return nums1[index-1], true
	}
	l1, l2 := 0, 0
	for {
		if l1 >= len(nums1) || l2 >= len(nums2) {
			break
		}
		if l1+l2+1 == index {
			return min(nums1[l1], nums2[l2]), true
		}
		if nums1[l1] < nums2[l2] {
			l1++
		} else {
			l2++
		}
	}
	if l1 == len(nums1) {
		return nums2[index-len(nums1)-1], true
	}
	if l2 == len(nums2) {
		return nums1[index-len(nums2)-1], true
	}
	return -1, false
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}

func min(num1, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}
