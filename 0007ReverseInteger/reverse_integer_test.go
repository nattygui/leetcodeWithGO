package _007ReverseInteger

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type question struct {
	one int
	ans int
}

func Test_ok(t *testing.T) {
	ast := assert.New(t)

	questions := []question{
		{
			one: 123,
			ans: 321,
		},
		{
			one: -12345,
			ans: -54321,
		},
		{
			one: 1<<31,
			ans: 0,
		},
	}

	for _, q := range questions {
		arg, res := q.one, q.ans
		ast.Equal(res, reverseInteger(arg), "输入为: %v", arg)
	}
}
