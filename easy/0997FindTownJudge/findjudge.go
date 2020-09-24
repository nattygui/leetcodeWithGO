package findjudge

type people struct {
	trusted  int
	trusting int
}

func findJudge(N int, trust [][]int) int {
	peopleMap := make(map[int]*people)

	for i := 1; i <= N; i++ {
		peopleMap[i] = &people{
			trusted:  0,
			trusting: 0,
		}
	}

	for _, t := range trust {
		peopleMap[t[0]].trusting++
		peopleMap[t[1]].trusted++
	}

	for i, p := range peopleMap {
		if p.trusted == N-1 && p.trusting == 0 {
			return i
		}
	}
	return -1
}
