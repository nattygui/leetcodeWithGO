package findjudge

func findJudge(N int, trust [][]int) int {
	inDegree := make([]int, N)
	outDegree := make([]int, N)
	for _, p := range trust {
		inDegree[p[1]-1]++
		outDegree[p[0]-1]++
	}
	result := -1
	for i := 0; i < N; i++ {
		if inDegree[i] == N-1 && outDegree[i] == 0 {
			result = i + 1
			break
		}
	}
	return result
}
