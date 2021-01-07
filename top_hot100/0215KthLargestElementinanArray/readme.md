# 215. 数组中的第K个最大元素
在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

示例 1:
```
输入: [3,2,1,5,6,4] 和 k = 2
输出: 5
```
示例 2:
```
输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
输出: 4
```

## 解法

借助 partition 操作定位到最终排定以后索引为 len - k 那个元素

```go
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
```