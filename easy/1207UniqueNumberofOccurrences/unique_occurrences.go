package uniqueoccurrences

func uniqueOccurrences(arr []int) bool {
	numsMap := make(map[int]int)
	for _, num := range arr {
		numsMap[num]++
	}

	numsExistMap := make(map[int]bool)
	for _, value := range numsMap {
		if numsExistMap[value] {
			return false
		}
		numsExistMap[value] = true
	}
	return true
}
