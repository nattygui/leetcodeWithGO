package isvalid

func isValid(s string) bool {
	// 构建括号map
	bracketsMap := map[byte]byte{
		'{': '}',
		'[': ']',
		'(': ')',
	}
	leftBrackets := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		switch s[i] {
		// 若是左括号，则加入切片中
		case '{', '[', '(':
			leftBrackets = append(leftBrackets, s[i])
		// 否则为右括号，通过括号map查找 切片最后一个元素是否匹配
		default:
			length := len(leftBrackets) - 1
			// 当前切片未含有有任何左括号，直接返回false
			if length == -1 {
				return false
			}
			// 若当前匹配，则删除最后一个，进行下一个
			if bracketsMap[leftBrackets[length]] == s[i] {
				leftBrackets = leftBrackets[:length]
			} else {
				return false // 不匹配，则直接返回false
			}
		}
	}
	if len(leftBrackets) != 0 { // 当遍历完全后，若当前还存在元素，则返回false
		return false
	}
	return true // 说明所有括号都匹配了，返回true
}
