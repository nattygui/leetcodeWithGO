# 33. 搜索旋转排序数组

时间复杂度：O(log n)

## 题解

通过查找当前 **mid** 为分割位置分割出来的俩个部分 **[1, mid]** 和 **[mid+1, r]** 哪个部分是有序的，并根据有序的那个部分确定我们如何改变二分搜索的上下届。

- 如果 **[l, mid - 1]** 是有序数组，且 target 的大小满足 **[nums[l],nums[mid])**，则我们应该将搜索范围缩小至 [l, mid - 1]，否则在 [mid + 1, r] 中寻找。

- 如果 **[mid, r]** 是有序数组，且 target 的大小满足 **(nums[mid+1],nums[r]]**，则我们应该将搜索范围缩小至 **[mid + 1, r]**，否则在 **[l, mid - 1]** 中寻找

## code 

```go

func search(nums []int, target int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}
	if length == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}

	l := 0
	r := length - 1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[0] <= nums[mid] {
			if nums[0] <= target && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[length-1] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}

	return -1
}
```