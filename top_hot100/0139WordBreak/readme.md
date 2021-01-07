# 139. 单词拆分
给定一个非空字符串 s 和一个包含非空单词的列表 wordDict，判定 s 是否可以被空格拆分为一个或多个在字典中出现的单词。

说明：

拆分时可以重复使用字典中的单词。
你可以假设字典中没有重复的单词。
示例 1：

```
输入: s = "leetcode", wordDict = ["leet", "code"]
输出: true
解释: 返回 true 因为 "leetcode" 可以被拆分成 "leet code"。
```

示例 2：

```
输入: s = "applepenapple", wordDict = ["apple", "pen"]
输出: true
解释: 返回 true 因为 "applepenapple" 可以被拆分成 "apple pen apple"。
     注意你可以重复使用字典中的单词。
```

示例 3：

```
输入: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
输出: false
```

## 解法一 动态规划

| "" | l | e | e | t | c | o | d | e |
| - | - | - | - | - | - | - | - | - |
| true | false | false | false | true | false | false | false | true |

初始化 dp=[true, False, ... ,False]，长度为 n+1。n 为字符串长度。dp[i] 表示 s 的前 i 位是否可以用 wordDictwordDict 中的单词表示。

初始化 dp[0]=true，空字符可以被表示。

遍历字符串的所有子串，遍历开始索引 ii，遍历区间 [0,n)：

遍历结束索引 j，遍历区间 [i+1,n+1)：
若 dp[i]=true 且 s[i,⋯,j) 在 wordlistwordlist 中：dp[j]=True。解释：dp[i]=True 说明 s 的前 i 位可以用 wordDictwordDict 表示，则 s[i,⋯,j) 出现在 wordDictwordDict 中，说明 ss 的前 j 位可以表示。
返回 dp[n]

复杂度分析
- 时间复杂度：O(n^2)
- 空间复杂度：O(n)

```go
func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	minLength := math.MaxInt64
	maxLength := math.MinInt64
	numWordsMap := make(map[int][]string)
	// 构建数组map
	for _, word := range wordDict {
		length := len(word)
		if minLength > length {
			minLength = length
		}
		if maxLength < length {
			maxLength = length
		}
		if v, ok := numWordsMap[length]; !ok {
			numWordsMap[length] = []string{word}
		} else {
			numWordsMap[length] = append(v, word)
		}
	}

	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			wordLen := j - i + 1
			if wordLen < minLength || wordLen > maxLength {
				continue
			}
			if words, ok := numWordsMap[j-i+1]; ok {
				word := s[i : j+1]
				for _, temp := range words {
					if dp[i] && temp == word {
						dp[j+1] = true
					}
				}
			}
		}
	}
	return dp[len(s)]
}
```

## 解法二 回溯法

golang
```go
func wordBreak(s string, wordDict []string) bool {
	numWordsMap := make(map[int][]string)
	// 构建数组map
	for _, word := range wordDict {
		length := len(word)
		if v, ok := numWordsMap[length]; !ok {
			numWordsMap[length] = []string{word}
		} else {
			numWordsMap[length] = append(v, word)
		}
	}

	return backtrack(s, numWordsMap)
}

func backtrack(s string, numWordsMap map[int][]string) bool {
	if len(s) == 0 {
		return true
	}
	result := false
	for i := 0; i < len(s); i++ {
		if words, ok := numWordsMap[i+1]; ok {
			word := s[:i+1]
			for _, temp := range words {
				if word == temp {
					result = backtrack(s[i+1:], numWordsMap) || result
                    break
				}
			}
		}
        if result {
            return result
        }
	}
	return result
}
```