package _190ReverseBits

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type question struct {
	one uint32
	ans uint32
}

func TestOK(t *testing.T) {
	ast := assert.New(t)

	questions := []question {
		{
			one: 43261596,
			ans: 964176192,
		},
	}

	for _, q := range questions {
		ast.Equal(q.ans, reverseBits(q.one))
	}

}
