# 0015 三数之和

给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。
 

示例：

```
给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]
``` 

## 解法一

在 0001俩数之和的基础上使用hashmap实现，但是数组长度过大的话，直接索引会造成较大的耗时

1. 首先先进行排序和构建全数组的hashmap，key为数值 value为索引

2. 然后进行俩层遍历，第二层的查找与题目 **0001俩数之和** 的实现相同，具体注释在代码中

- 时间复杂度 O(n^2)

- 空间复杂度 O(n)

但是在数据量比较大的情况下，通过分析pprof中由于map key的索引比较耗时，所以整体会比较耗时

### code

```go
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	// 构建numsMap
	numsMap := make(map[int]int)
	for index, num := range nums {
		numsMap[num] = index
	}
	// 开始遍历
	result := make([][]int, 0)
	for index := 0; index < len(nums)-2; index++ { // 第一个数值
		if index > 0 && nums[index] == nums[index-1] { // 过滤重复值
			continue
		}

		for index1 := index + 1; index1 < len(nums)-1; index1++ { // 查找第二个数值
			if index1 > index+1 && nums[index1] == nums[index1-1] { // 过滤重复值
				continue
			}
			if v, ok := numsMap[-nums[index]-nums[index1]]; ok && index1 < v { // 查找第三个数值，若当前元素在map中，且索引大于第二个数值的索引，则为找到
				result = append(result, []int{nums[index], nums[index1], -nums[index] - nums[index1]}) 
			}
		}
	}
	return result
}
```

## 解法二 

与解法一大致相同，在查找第二个数值和第三个数值的时候使用的二分查找

- 时间复杂度 O(n^2)

- 空间复杂度 O(1)

```go
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	length := len(nums)
	// 开始遍历
	var start, end int
	result := make([][]int, 0)
	for index := 0; index < length-2; index++ {
		if index > 0 && nums[index] == nums[index-1] {
			continue
		}
		start = index + 1
		end = length - 1
		for start < end {
			if start > index+1 && nums[start] == nums[start-1] {
				start++
				continue
			}
			if end < length-2 && nums[end] == nums[end+1] {
				end--
				continue
			}
			if nums[index]+nums[start]+nums[end] > 0 {
				end--
			} else if nums[index]+nums[start]+nums[end] < 0 {
				start++
			} else {
				result = append(result, []int{nums[index], nums[start], nums[end]})
				end--
				start++
			}
		}
	}
	return result
}
```