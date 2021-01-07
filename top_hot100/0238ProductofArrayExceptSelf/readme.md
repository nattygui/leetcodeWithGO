# 238. 除自身以外数组的乘积
给你一个长度为 n 的整数数组 nums，其中 n > 1，返回输出数组 output ，其中 output[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积。

示例:
```
输入: [1,2,3,4]
输出: [24,12,8,6]
```

## 解法

```go
func productExceptSelf(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	if len(nums) == 1 {
		return []int{1}
	}
	answer := make([]int, len(nums))
	answer[0] = 1
	for i := 1; i < len(nums); i++ {
		answer[i] = answer[i-1] * nums[i-1]
	}

	r := 1
	for i := len(nums) - 2; i >= 0; i-- {
        r *= nums[i+1]
		answer[i] *= r
	}
	return answer
}
```