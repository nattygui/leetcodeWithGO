# 17. 电话号码的字母组合

## 解法

回溯法：

```go
func letterCombinations(digits string) []string {
	// 如果没有输入，则返回空
	if len(digits) == 0 {
		return []string{}
	}
	// 1. 首先构建一个长度为8的[]string 数组
	// 分别为 索引 0 对应数字为 2 对应字符串数组 {"a", "b", "c"}
	// 索引 1 对应数字为 3 对应字符串数组 {"d", "e", "f"}
	// 索引 2 对应数字为 4 对应字符串数组 {"g", "h", "i"}
	// 索引 3 对应数字为 5 对应字符串数组 {"j", "k", "l"}
	// 索引 4 对应数字为 6 对应字符串数组 {"m", "n", "o"}
	// 索引 5 对应数字为 7 对应字符串数组 {"p", "q", "r", "s"}
	// 索引 6 对应数字为 8 对应字符串数组 {"t", "u", "v"}
	// 索引 7 对应数字为 9 对应字符串数组 {"w", "x", "y", "z"}
	letter := byte('a')
	numbers := make([][]string, 8)
	for i := 0; i < len(numbers); i++ {
		num := 3
		if i == 5 || i == 7 {
			num = 4
		}
		strSlice := make([]string, num)
		for j := 0; j < num; j++ {
			strSlice[j] = string(letter)
			letter++
		}
		numbers[i] = strSlice
	}

	// 确定返回的结果的大小，减少多次内存分配
	allStrNums := 0
	for index := range digits {
		allStrNums *= len(numbers[digits[index]-byte('2')])
	}

	result := make([]string, 0, allStrNums)
	// 开始逐层扫描
	matchAllLetter(0, digits, numbers, "", &result)
	return result
}

func matchAllLetter(level int, digits string, numbers [][]string, curResult string, result *[]string) {
	digit := int(digits[level] - byte('2'))
	// 遍历当前层的所有字符
	for _, curNum := range numbers[digit] {
		temp := curResult + curNum
		// 若是最后一个字符，则加入结果中，否则将当前当前层的结果传入下一层
		if level == len(digits)-1 {
			*result = append(*result, temp)
		} else {
			matchAllLetter(level+1, digits, numbers, temp, result)
		}
	}
}
```