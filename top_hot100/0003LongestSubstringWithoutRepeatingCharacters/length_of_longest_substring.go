package lengthoflongestsubstring

func lengthOfLongestSubstring(s string) int {
	charsFreq := make([]int, 128) // 记录最新出现的字符出现的位置
	for index := range charsFreq {
		charsFreq[index] = -1
	}
	flag := 0   // 标记位
	result := 0 // 结果
	for index := 0; index < len(s); index++ {
		preIndex := charsFreq[s[index]]
		if flag > preIndex { // 若当前flag 大于之前的出现的位置
			result = max(result, index-flag+1)
		} else { // 若当前flag 小于之前出现的位置
			result = max(result, index-flag)
			flag = preIndex + 1 // 替换最新的位置
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
