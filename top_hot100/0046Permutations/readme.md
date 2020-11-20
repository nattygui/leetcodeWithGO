# 46. 全排列

给定一个 没有重复 数字的序列，返回其所有可能的全排列。

示例:

```
输入: [1,2,3]
输出:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]
```

## 解法

多层递归，查找可能存在的值

```go
func permute(nums []int) [][]int {
	result := make([][]int, 0)
	walk(nums, nil, &result, len(nums))
	return result
}

func walk(nums []int, oneResult []int, result *[][]int, length int) {
	if oneResult == nil {
		oneResult = make([]int, 0, length)
	}
	for i, num := range nums {
		tempResult := append(oneResult, num)
		nextNums := make([]int, 0, len(nums)-1)
		nextNums = append(nextNums, nums[:i]...)
		nextNums = append(nextNums, nums[i+1:]...)
		if len(tempResult) == length {
			// 添加一个结果的时候，由于slice 的特性，需要重新创建一个切片
			temp := make([]int, length)
			copy(temp, tempResult)
			*result = append(*result, temp)
		} else {
			walk(nextNums, tempResult, result, length)
		}
	}
}
```
