package _575DistributeCandies

func distributeCandies(candies []int) int {
	kv := make(map[int]int)

	for _, v := range candies {
		if _, ok := kv[v]; ok {
			kv[v]++
		} else {
			kv[v] = 0
		}
	}

	if len(candies) / 2 > len(kv) {
		return len(kv)
	}
	return len(candies) / 2
}