package _118PascalsTriangle

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type question struct {
	one int
	res [][]int
}

func TestOK(t *testing.T) {
	ast := assert.New(t)

	questions := []question{
		{
			one: 5,
			res: [][]int{
				{1},
				{1, 1},
				{1, 2, 1},
				{1, 3, 3, 1},
				{1, 4, 6, 4, 1},
			},
		},
	}

	for _, q := range questions {
		ast.Equal(q.res, generate(q.one), "输入为%v", q.one)
	}
}