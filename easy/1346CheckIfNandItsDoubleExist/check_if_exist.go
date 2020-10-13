package checkifexist

func checkIfExist(arr []int) bool {
	numMaps := make(map[int]int, 0)
	for index, num := range arr {
		numMaps[num] = index
	}

	for index, num := range arr {
		if value, ok := numMaps[num*2]; ok && index != value {
			return true
		}
	}
	return false
}
