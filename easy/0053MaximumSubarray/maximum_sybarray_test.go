package _053MaximumSubarray

import (
	"github.com/stretchr/testify/assert"
	"testing"
)



type question struct {
	one []int
	res int
}

func TestOK(t *testing.T) {
	ast := assert.New(t)

	questions := []question{
		{
			one: []int{-2,1,-3,4,-1,2,1,-5,4},
			res: 6,
		},
	}

	for _, q := range questions {
		ast.Equal(q.res, maxSubArray0(q.one), "输入为: %v", q.one)
		ast.Equal(q.res, maxSubArray1(q.one), "输入为: %v", q.one)
		ast.Equal(q.res, maxSubArray2(q.one), "输入为: %v", q.one)
	}
}
