package _205IsomorphicStrings

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type question struct {
	one string
	two string
	ans bool
}

func TestOK(t *testing.T) {
	ast := assert.New(t)

	questions := []question{
		{
			one: "eqq",
			two: "tww",
			ans: true,
		},
	}

	for _, q := range questions {
		ast.Equal(q.ans, isIsomorphic(q.one, q.two))
	}

}
