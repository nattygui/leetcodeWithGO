package mostvisited

func mostVisited(n int, rounds []int) []int {
	results := make([]int, n)
	for i := 1; i < len(rounds); i++ {
		start := rounds[i-1]
		end := rounds[i]
        addSteps(start, end, n, &results)
	}
	return results
}

func addSteps(start, end, n int, results *[]int) {
    if start < end {
        for start < end {
            start++
            (*results)[start-1]++
        }
    } else {
        for start < n {
            start++
            (*results)[start-1]++
        }
        start = 0
        for start < end {
            start++
            (*results)[start-1]++
        }
    }
}
