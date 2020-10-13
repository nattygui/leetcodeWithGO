package kweakestrows

import "sort"

func kWeakestRows(mat [][]int, k int) []int {
	length := len(mat)
	indies := make([]int, length)
	for i := 0; i < length; i++ {
		indies[i] = i + 1
	}

	scores := make([]int, length)
	for i := 0; i < length; i++ {
		num := 0
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 1 {
				num++
			}
		}
		scores[i] = num
	}

	sr := &scoreResult{
		indies: indies,
		scores: scores,
	}
	sort.Sort(sr)
	return sr.indies[:k]
}

type scoreResult struct {
	indies []int
	scores []int
}

func (sr *scoreResult) Len() int { return len(sr.scores) }
func (sr *scoreResult) Swap(i, j int) {
	sr.indies[i], sr.indies[j] = sr.indies[j], sr.indies[i]
	sr.scores[i], sr.scores[j] = sr.scores[j], sr.scores[i]
}
func (sr *scoreResult) Less(i, j int) bool {
	if sr.scores[i] == sr.scores[j] {
		return sr.indies[i] < sr.indies[j]
	}
	return sr.scores[i] < sr.scores[j]
}
