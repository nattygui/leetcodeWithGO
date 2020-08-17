package _122BestTimeToBuyAndSellStockII

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
			ans: 7,
		},
		{
			one: []int{7, 6, 4, 3, 1},
			ans: 0,
		},
	}

	for _, q := range questions {
		ast.Equal(q.ans, maxProfitOne(q.one), "输入为%v", q.one)
	}
}
