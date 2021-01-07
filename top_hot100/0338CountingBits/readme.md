# 338. 比特位计数

给定一个非负整数 num。对于 0 ≤ i ≤ num 范围中的每个数字 i ，计算其二进制数中的 1 的数目并将它们作为数组返回。

示例 1:
```
输入: 2
输出: [0,1,1]
```
示例 2:
```
输入: 5
输出: [0,1,1,2,1,2]
```

## 解法

```go
func countBits(num int) []int {
	dp := make([]int, num+1)
	for i := 0; i <= num; i++ {
		if i == 0 {
			dp[i] = 0
		} else if i == 1 || i == 2 {
			dp[i] = 1
		} else {
			dp[i] = findValue(dp, i)
		}
	}
	return dp
}

func findValue(dp []int, i int) int {
	if i <= 2 {
		return dp[i]
	}
	if i/2 == 0 {
		return dp[i/2]
	}
	return dp[i-1] + 1
}
```