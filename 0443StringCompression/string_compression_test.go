package _443StringCompression

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type question struct {
	one []byte
	ans int
}

func TestOk(t *testing.T) {
	ast := assert.New(t)

	questions := []question{
		{
			one: []byte{'a', 'b', 'b', 'c', 'c', 'c', 'c'},
			ans: 5,
		},
	}

	for _, q := range questions {
		ast.Equal(q.ans, compress(q.one))
	}
}
