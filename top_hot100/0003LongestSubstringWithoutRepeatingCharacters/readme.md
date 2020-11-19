# 0003. 无重复字符的最长子串

给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

```
输入: "abcabcbb"
输出: 3 
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
```

示例 2:

```
输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
```

示例 3:

```
输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
```

## 解法一  用来标记位置

查找无重复的最大子字符串，其实就是处理好边界问题。

1. 设置一个flag位，用来定位最新的索引起点。而字符串切片 用来记录当前字符最新出现的位置

2. 遍历当前字符串 有以下俩种情况。

    当前字符索引值为index，当前查找的字符s[index]在字符串切片中存储的位置为 preindex

    - 若当前字符未出现 或者 出现过但是flag大于当前字符的位置，则直接判断 max < index - flag - 1 若小于则 max = index - flag + 1

    - 若当前字符之前出现且flag 小于preIndex则, 则直接判断 max < index - flag 若小于 则max = index - flag, 并替换当前flag 位置 flag = preIndex + 1

- 时间复杂度:   O(n)

- 空间复杂度:   O(n)

## code

执行用时：
0 ms, 在所有 Go 提交中击败了100.00%的用户

内存消耗：
2.5 MB, 在所有 Go 提交中击败了89.80%的用户

```go
func lengthOfLongestSubstring(s string) int {
	charsFreq := make([]int, 128) // 记录最新出现的字符出现的位置
	for index := range charsFreq {
		charsFreq[index] = -1
	}
	flag := 0   // 标记位，起点位
	result := 0 // 结果
	for index := 0; index < len(s); index++ {
		preIndex := charsFreq[s[index]]
		if flag > preIndex { // 若当前flag 大于之前的出现的位置
			result = max(result, index-flag+1)
		} else { // 若当前flag 小于之前出现的位置
			result = max(result, index-flag)
			flag = preIndex + 1 // 替换最新标记位，也是起点位
		}
		charsFreq[s[index]] = index // 重新记录当前字符位置
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
