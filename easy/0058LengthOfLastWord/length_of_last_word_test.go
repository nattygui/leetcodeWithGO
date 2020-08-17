package _058LengthOfLastWord

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type question struct {
	one string
	res int
}

func TestOK(t *testing.T) {
	ast := assert.New(t)

	questions := []question{
		{
			one: "Hello World",
			res: 5,
		},
		{
			one: "Hello go",
			res: 2,
		},
		{
			one: " goland",
			res: 6,
		},
	}

	for _, q := range questions {
		ast.Equal(q.res, lengthOfLastWord(q.one), "输入为：%v", q.one)
	}
}