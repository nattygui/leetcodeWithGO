# 0005. 最长回文子串

给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：

```
输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
```

示例 2：

```
输入: "cbbd"
输出: "bb"
```

## 解法一（暴力解法）

1. 通过建立双重索引，索引i从0开始，索引j从最后一个元素开始，通过辅助函数s[i:j+1] 来判断当前范围是否是题解。

- 时间复杂度: O(n^3)
- 空间复杂度: O(1)

### code

```go
func longestPalindrome(s string) string {
	start, end := 0, 0
	max := 0
	for i := 0; i < len(s); i++ {
		for j := len(s) - 1; j >= i; j-- {
			if isPalindrome(s[i:j+1]) && j+1-i > max {
				start = i
				end = j + 1
				max = end - start
			}
		}
	}
	return s[start:end]
}

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
```

## 解法二 （动态规划)

- 时间复杂度: O(n^2)

- 空间复杂度: O(n^2)

```go
func longestPalindrome(s string) string {
	length := len(s)
	result := ""
	dp := make([][]int, length) // 构建nxn的数组
	for index := range dp {
		dp[index] = make([]int, length)
	}
	for l := 0; l < length; l++ {
		for i := 0; i+l < length; i++ {
			j := i + l
			if l == 0 { // 为0的时候进行初始化
				dp[i][j] = 1
			} else if l == 1 { // 为1的时候进行初始化
				if s[i] == s[j] {
					dp[i][j] = 1
				}
			} else {
				if s[i] == s[j] { // 状态转移
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] > 0 && l+1 > len(result) {
				result = s[i : j+1] // 获取最优的结果
			}
		}
	}
	return result
}
```

## 解法三（中心拓展法）

- 时间复杂度: O(n^2)

- 空间复杂度: O(1)

```go
func longestPalindrome(s string) string {
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		start1, end1 := findMostResult(s, i, i)
		start2, end2 := findMostResult(s, i, i+1)
		if end1-start1 > end-start {
			end = end1
			start = start1
		}
		if end2-start2 > end-start {
			end = end2
			start = start2
		}
	}
	return s[start : end+1]
}

func findMostResult(s string, i, j int) (int, int) {
	for ; i >= 0 && j < len(s) && s[i] == s[j]; i, j = i-1, j+1 {
	}
	return i + 1, j - 1
}
```