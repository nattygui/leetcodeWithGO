package _026RemoveDuplicatesFromSortedArray

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type question struct {
	one []int
	res int
}

func Test_ok(t *testing.T) {
	ast := assert.New(t)

	questions := []question{
		{
			one: []int{0,0,1,1,1,2,2,3,3,4},
			res: 5,
		},
		{
			one: []int{},
			res: 0,
		},
	}

	for _, q := range questions {
		ast.Equal(q.res, removeDuplicates(q.one), "输入为:%v", q.one)
	}
}