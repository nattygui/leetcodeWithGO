package _447NumberOfBoomerangs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type question struct {
	one [][]int
	ans int
}

func TestOK(t *testing.T) {
	ast := assert.New(t)

	q := question{
		one: [][]int{{0,0},{1,0},{2,0}},
		ans: 2,
	}

	ast.Equal(q.ans, numberOfBoomerangs(q.one))

}

