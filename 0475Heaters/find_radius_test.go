package _475Heaters

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type question struct {
	one []int
	two []int
	ans int
}

func TestOk(t *testing.T) {
	ast := assert.New(t)

	q := question{
		one: []int{2, 4, 5},
		two: []int{2,4,5},
		ans: 1,
	}

	ast.Equal(q.ans, findRadiusTwo(q.one, q.two))
}