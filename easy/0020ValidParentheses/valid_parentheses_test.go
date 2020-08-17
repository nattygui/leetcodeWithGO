package _020ValidParentheses

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type question struct {
	one string
	ans bool
}

func Test_ok(t *testing.T) {
	ast := assert.New(t)

	questions := []question{
		{
			one: "()",
			ans: true,
		},
		{
			one: "()[]{}",
			ans: true,
		},
		{
			one: "(]",
			ans: false,
		},
		{
			one: "([)]",
			ans: false,
		},
		{
			one: "{[]}",
			ans: true,
		},
	}

	for _, q := range questions {
		ast.Equal(q.ans, isValid(q.one), "输入是：%v", q.one)
	}
}
