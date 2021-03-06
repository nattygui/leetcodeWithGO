package groupanagrams

func groupAnagrams(strs []string) [][]string {
	// 最终返回的结果
	result := make([][]string, 0)

	// num 用于记录当前出现不同词的数量
	num := 0
	// map 记录出现过的字符
	// key 为记录字符切片， value 为result的索引数组
	strMap := make(map[[26]byte]int)
	for _, str := range strs {
		// 计算当前字符
		strBytes := [26]byte{}
		for i := 0; i < len(str); i++ {
			strBytes[str[i]-'a']++
		}
		// 若当前map 不存在，则添加到map中， num为第n个不同小写字母组成的词
		if index, ok := strMap[strBytes]; !ok {
			result = append(result, []string{str})
			strMap[strBytes] = num
			num++
		} else { // 若存在，根据存储的第n个列，将当前str 添加进去
			result[index] = append(result[index], str)
		}
	}
	return result
}
