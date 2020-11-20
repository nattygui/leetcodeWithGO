# 53. 最大子序和

给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

示例:

```
输入: [-2,1,-3,4,-1,2,1,-5,4]
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
```

# 解法- 贪心算法

若当前数组值 nums[i] + prev < nums[i], 说明prev 对子序列和没有贡献，我们就丢弃当前元素之前的数列

```go 
func maxSubArray(nums []int) int {
	result := nums[0]
	prev := nums[0]
	for i := 1; i < len(nums); i++ {
		prev = max(nums[i], nums[i]+prev)
		result = max(prev, result)
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

# 解法二 动态规划

状态转移方程： 1. if dp[i-1] > 0, dp[i] = nums[i]+ dp[i-1] ; else dp[i] =  nums[i]

```go
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			nums[i] = nums[i] + nums[i-1]
		}
		if max < nums[i] {
			max = nums[i]
		}
	}

	return max
}
```
