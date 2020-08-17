package _682BaseballGame

import "strconv"

type Result struct {
	scores []int
	res    int
	len    int
}

func (r *Result) retriveScore(s string) {
	switch s {
	case "C":
		temp := r.scores[r.len]
		r.res -= temp
		r.scores = r.scores[0:r.len]
		r.len--
	case "D":
		temp := r.scores[r.len] * 2
		r.scores = append(r.scores, temp)
		r.len++
		r.res += temp
	case "+":
		r.res += r.scores[len(r.scores)-1] + r.scores[len(r.scores)-2]
		r.scores = append(r.scores, r.scores[len(r.scores)-1]+r.scores[len(r.scores)-2])
		r.len++
	default:
		temp, _ := strconv.Atoi(s)
		r.scores = append(r.scores, temp)
		r.res += temp
		r.len++
	}
}

func calPoints(ops []string) int {
	result := &Result{
		scores: []int{},
		res:    0,
		len:    -1,
	}
	for _, str := range ops {
		result.retriveScore(str)
	}
	return result.res
}
