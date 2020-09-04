package numspecialequivgroups

type stringsNode struct {
	oddBytes  map[byte]int
	evenBytes map[byte]int
}

func numSpecialEquivGroups(A []string) int {
	stringsGroup := make([]stringsNode, 0)

	for _, str := range A {
		oddTemp := make(map[byte]int)
		evenTemp := make(map[byte]int)
		for i := 0; i < len(str); i++ {
			if i%2 == 0 {
				oddTemp[str[i]]++
			} else {
				evenTemp[str[i]]++
			}
		}
		stringsGroup = append(stringsGroup, stringsNode{oddBytes: oddTemp, evenBytes: evenTemp})
	}

	return splitSlice(stringsGroup)
}

func splitSlice(nodes []stringsNode) int {
	nodeMap := make(map[string]int)
	for _, node := range nodes {
		str := node.toString()
		nodeMap[str]++
	}

	return len(nodeMap)
}

func (s stringsNode) toString() string {
	oddLetters := make([]int, 26)
	evenLetters := make([]int, 26)

	res := make([]byte, 0)

	for key, value := range s.oddBytes {
		oddLetters[int(key-'a')] += value
	}

	for key, value := range s.evenBytes {
		evenLetters[int(key-'a')] += value
	}

	for key, value := range oddLetters {
		res = append(res, byte(key))
		res = append(res, byte(value))
	}

	for key, value := range evenLetters {
		res = append(res, byte(key))
		res = append(res, byte(value))
	}
	return string(res)
}
