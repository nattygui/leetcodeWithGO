package _121BestTimeToBuyAndSellStock

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type question struct {
	one []int
	ans int
}

func TestOK(t *testing.T) {
	ast := assert.New(t)

	questions := []question{
		{
			one: []int{7, 1, 5, 3, 6, 4},
			ans: 5,
		},
		{
			one: []int{7,6,4,3,1},
			ans: 0,
		},
	}

	for _, q := range questions {
		ast.Equal(q.ans, maxProfitTwo(q.one), "输入为%v", q.one)
	}
}