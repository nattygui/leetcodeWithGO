# 300. 最长上升子序列
给定一个无序的整数数组，找到其中最长上升子序列的长度。

示例:
```
输入: [10,9,2,5,3,7,101,18]
输出: 4 
解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。
```

## 解法一 暴力解法

逐个查找当前数的存在比较大的值

```go
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
        return 0
    }
    result := 1
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > nums[i] {
				findMaxLength(nums, j, 2, &result)
			}
		}
	}
	return result
}

func findMaxLength(nums []int, x, count int, result *int) {
	if count > *result {
        *result = count
    }
    for i := x + 1; i < len(nums); i++ {
		if nums[i] > nums[x] {
			temp := count+1
			findMaxLength(nums, i, temp, result)
		}
	}
}
```

## 解法二 动态规划

```go
func lengthOfLIS(nums []int) int {
	result := 0
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
        dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		if result < dp[i] {
			result = dp[i]
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
```

