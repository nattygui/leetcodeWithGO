package _088MergeSortedArray

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type param struct {
	one []int
	two int
	three []int
	four int
}

type question struct {
	params param
	res []int
}

func TestOK(t *testing.T) {
	ast := assert.New(t)

	questions := []question{
		{
			params: param{
				one:   []int{1, 2, 3, 0, 0, 0},
				two:   3,
				three: []int{0, 0, 0},
				four:  3,
			},
			res: []int{0, 0, 0, 1, 2, 3},
		},
	}
	for _, q := range questions {
		ast.Equal(q.res, merge(q.params.one, q.params.two, q.params.three, q.params.four), "输入为: %v", q.params)
	}
}