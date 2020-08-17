package _066PlusOne

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type question struct {
	one []int
	res []int
}

func TestOK(t *testing.T) {
	ast := assert.New(t)

	questions := []question{
		{
			one: []int{1,2,3,4,5},
			res: []int{1,2,3,4,6},
		},
		{
			one: []int{1,2,9,9,9},
			res: []int{1,3,0,0,0},
		},
		{
			one: []int{9,9,9,9,9},
			res: []int{1,0,0,0,0,0},
		},
	}
	for _, q := range questions {
		ast.Equal(q.res, plusOne(q.one), "输入为: %v", q.one)
	}
}
