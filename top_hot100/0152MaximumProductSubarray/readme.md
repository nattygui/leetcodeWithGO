# 152. 乘积最大子数组
给你一个整数数组 nums ，请你找出数组中乘积最大的连续子数组（该子数组中至少包含一个数字），并返回该子数组所对应的乘积。

示例 1:

```
输入: [2,3,-2,4]
输出: 6
解释: 子数组 [2,3] 有最大乘积 6。
```

示例 2:

```
输入: [-2,0,-1]
输出: 0
解释: 结果不能为 2, 因为 [-2,-1] 不是子数组。
```

## 解法

golang
```go
func maxProduct(nums []int) int {
	result := nums[0]
	minDp := make([]int, len(nums))
	maxDp := make([]int, len(nums))
	for index, num := range nums {
		if index == 0 {
			minDp[0] = num
			maxDp[0] = num
			continue
		}
		minDp[index] = min(minDp[index-1]*num, min(maxDp[index-1]*num, num))
		maxDp[index] = max(maxDp[index-1]*num, max(minDp[index-1]*num, num))
		if result < maxDp[index] {
			result = maxDp[index]
		}
		if result < minDp[index] {
			result = minDp[index]
		}
	}
	return result
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
```