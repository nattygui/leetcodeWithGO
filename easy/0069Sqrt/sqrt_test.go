package _069Sqrt

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
			one: 5,
			res: 2,
		},
		{
			one: 25,
			res: 5,
		},
	}

	for _, q := range questions {
		ast.Equal(q.res, mySqrt(q.one), "输入为: %v", q.one)
	}
}
