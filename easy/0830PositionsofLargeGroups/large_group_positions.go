package largegrouppositions

func largeGroupPositions(S string) [][]int {
	res := make([][]int, 0)

	start := 0
	end := 0
	var char byte
	for i := 0; i < len(S); i++ {
		if char == S[i] {
			end++
			if i == len(S)-1 {
				if end-start >= 2 {
					res = append(res, []int{start, end})
				}
				return res
			}
			continue
		}

		if end-start >= 2 {
			res = append(res, []int{start, end})
		}
		start = i
		end = i
		char = S[i]
	}
	return res
}
