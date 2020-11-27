# 96. 不同的二叉搜索树
给定一个整数 n，求以 1 ... n 为节点组成的二叉搜索树有多少种？

示例:

输入: 3
输出: 5
解释:
给定 n = 3, 一共有 5 种不同结构的二叉搜索树:

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3

## 解决

golang
```go
func numTrees(n int) int {
    if n <= 0 {
        return 0
    }
	dp := make([]int, n+1)
    dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		all := 0
		for j := 1; j <= i; j++ {
			all += dp[j-1] * dp[i-j]
		}
		dp[i] = all
	}
	return dp[n]
}
```

python
```py
class Solution:
    def numTrees(self, n):
        dp = [1, 1]
        for i in range(1, n+1):
            nums = 0
            for j in range(1, i):
                nums += dp[j-1] * dp[i-j]
            dp.append(nums)
        
        return dp[n]
```

