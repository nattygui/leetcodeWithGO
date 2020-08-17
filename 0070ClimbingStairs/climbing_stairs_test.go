package _070ClimbingStairs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type question struct {
	one int
	res int
}

func TestOK(t *testing.T) {
	ast := assert.New(t)

	questions := []question{
		{
			one: 10,
			res: 89,
		},
		{
			one: 3,
			res: 3,
		},
	}

	for _, q := range questions {
		ast.Equal(q.res, climbStairs(q.one), "输入为: %v", q.one)
	}
}