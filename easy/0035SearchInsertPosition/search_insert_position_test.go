package _035SearchInsertPosition

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type param struct {
	one []int
	two int
}

type question struct {
	params param
	res int
}

func TestOK(t *testing.T) {
	ast := assert.New(t)

	questions := []question{
		{
			params: param{
				one: []int{1, 3, 5, 6},
				two: 5,
			},
			res: 2,
		},
		{
			params: param{
				one: []int{1, 3, 5, 6},
				two: 2,
			},
			res: 1,
		},
		{
			params: param{
				one: []int{1, 3, 5, 6},
				two: 7,
			},
			res: 4,
		},
		{
			params: param{
				one: []int{1, 3, 5, 6},
				two: 0,
			},
			res: 0,
		},
	}
	for _, q := range questions {
		ast.Equal(q.res, searchInsert(q.params.one, q.params.two), "输入为: %v", q.params)
	}
	for _, q := range questions {
		ast.Equal(q.res, searchInsertBinary(q.params.one, q.params.two), "输入为: %v", q.params)
	}
}
