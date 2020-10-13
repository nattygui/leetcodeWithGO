package findlucky

func findLucky(arr []int) int {
	res := -1
	staticMap := make(map[int]int)
	for _, num := range arr {
		staticMap[num]++
	}

	for key, value := range staticMap {
		if key == value && res < key {
			res = key
		}
	}
	return res
}
