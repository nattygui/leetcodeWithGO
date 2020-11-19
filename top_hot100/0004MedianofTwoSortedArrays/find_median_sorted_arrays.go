package findmediansortedarrays

import "sort"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	length1 := len(nums1)
	length2 := len(nums2)
	if length1 == 0 && length2 == 0 {
		return 0.0
	}
	leftTarget, rightTarget := 0, 0
	if (length1+length2)%2 == 0 {
		leftTarget, rightTarget = (length1+length2)/2-1, (length1+length2)/2
		return (float64(findKthElement(nums1, nums2, leftTarget)) + float64(findKthElement(nums1, nums2, rightTarget))) / 2
	}
	leftTarget = (length1 + length2) / 2
	return float64(findKthElement(nums1, nums2, leftTarget))
}

func findKthElement(nums1 []int, nums2 []int, indexTarget int) int {
	index1, index2 := 0, 0
	if len(nums1) == 0 {
		return nums2[indexTarget]
	}
	if len(nums2) == 0 {
		return nums1[indexTarget]
	}
	for {
		if index1+index2+1 == indexTarget {
			if index1 == len(nums1)-1 || index2 == len(nums2)-1 {
				return max(nums1[index1], nums2[index2])
			}
			return min2(nums1[index1], nums2[index2], nums1[index1+1], nums2[index2+1])
		}
		if nums1[index1] < nums2[index2] {
			if index1 != len(nums1)-1 {
				index1++
			} else {
				index2++
			}
		} else {
			if index2 != len(nums2)-1 {
				index2++
			} else {
				index1++
			}
		}
	}
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}

func min2(nums ...int) int {
	sort.Ints(nums)
	return nums[1]
}
