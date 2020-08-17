package _119PascalsTriangleII

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type question struct {
	one int
	ans []int
}

func TestOK(t *testing.T) {
	ast := assert.New(t)

	questions := []question{
		{
			one: 3,
			ans: []int{1, 3, 3, 1},
		},
		{
			one: 4,
			ans: []int{1, 4, 6, 4, 1},
		},
	}

	for _, q := range questions {
		ast.Equal(q.ans, getRowTwo(q.one), "输入为%v", q.one)
	}
}
