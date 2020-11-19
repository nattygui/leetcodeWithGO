# 31. 下一个列表

## code

```go
func nextPermutation(nums []int) {
	// 1 从倒序开始查找是否存在 之前的值小于 后面的值
	flag := false
	maxIndex := len(nums) - 1
	var startIndex int
	for j := len(nums) - 2; j >= 0; j-- {
		if nums[maxIndex] < nums[j] {
			maxIndex = j
		} else {
			startIndex = j
			flag = true
			break
		}
	}
	// 若存在
	if flag {
		// 找到最接近的index
		nearstIndex := 0
		num := nums[startIndex]
		diff := nums[maxIndex] - num
		for i := startIndex + 1; i < len(nums); i++ {
			if nums[i]-nums[startIndex] > 0 && diff > nums[i]-nums[startIndex] {
				diff = nums[i] - nums[startIndex]
				nearstIndex = i
			}
		}
		// 找到后交换 nearstIndex 和 startIndex
		nums[startIndex], nums[nearstIndex] = nums[nearstIndex], nums[startIndex]
		// 从startIndex+1 开始对之后的所有元素进行排序
		var quicksort func(start int, end int)
		quicksort = func(start int, end int) {
			i := start
			j := end
			minIndex := start
			for i < j {
				for i < j && nums[j] > nums[minIndex] {
					j--
				}
				if i < j {
					nums[i] = nums[j]
					i++
				}
				for i < j && nums[i] < nums[minIndex] {
					i++
				}
				if i < j {
					nums[j] = nums[i]
					j--
				}
			}
			nums[i] = nums[minIndex]
			quicksort(start, i-1)
			quicksort(i+1, end)
		}
		quicksort(startIndex, len(nums)-1)
	} else { // 若不存在，直接反转
		i := 0
		j := len(nums) - 1
		for i < j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
	}
}
```
