package _204CountPrimes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type question struct {
	one int
	ans int
}

func TestOK(t *testing.T) {
	ast := assert.New(t)

	questions := []question{
		{
			one: 10,
			ans: 4,
		},
	}

	for _, q := range questions {
		ast.Equal(q.ans, countPrimes(q.one))
	}
}