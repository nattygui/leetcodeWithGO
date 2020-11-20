package groupanagrams

func groupAnagrams(strs []string) [][]string {
	result := make([][]string, 0)
	num := 0
	strMap := make(map[[26]byte]int)
	for _, str := range strs {
		strBytes := [26]byte{}
		for i := 0; i < len(str); i++ {
			strBytes[str[i]-'a']++
		}
		if index, ok := strMap[strBytes]; !ok {
			result = append(result, []string{str})
			strMap[strBytes] = num
			num++
		} else {
			result[index] = append(result[index], str)
		}
	}
	return result
}
