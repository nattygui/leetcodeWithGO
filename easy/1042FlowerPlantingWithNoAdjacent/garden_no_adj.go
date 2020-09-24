package garden

func gardenNoAdj(N int, paths [][]int) []int {
	adjacent := make([][]int, N)
	for _, path := range paths {
		adjacent[path[0]-1] = append(adjacent[path[0]-1], path[1]-1)
		adjacent[path[1]-1] = append(adjacent[path[1]-1], path[0]-1)
	}

	result := make([]int, N)
	for i := 0; i < N; i++ {
		used := make([]bool, 4)
		for _, node := range adjacent[i] {
			if result[node] != 0 {
				used[result[node]-1] = true
				continue
			}
		}
		if result[i] == 0 {
			for index, ok := range used {
				if !ok {
					used[index] = true
					result[i] = index + 1
					break
				}
			}
		}
	}
	return result
}
