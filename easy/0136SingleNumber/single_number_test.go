package _136SingleNumber

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type question struct {
	one []int
	ans int
}

func TestOk(t *testing.T) {
	ast := assert.New(t)

	questions := []question{
		{
			one: []int{2, 2, 1},
			ans: 1,
		},
		{
			one: []int{4, 1, 2, 1, 2},
			ans: 4,
		},
	}

	for _, q := range questions {
		ast.Equal(q.ans, singleNumber(q.one), "输入为%v", q.one)
	}

}
