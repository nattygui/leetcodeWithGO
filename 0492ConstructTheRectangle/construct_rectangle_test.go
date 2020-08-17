package _492ConstructTheRectangle

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type question struct {
	one int
	res []int
}

func TestOk(t *testing.T) {
	ast := assert.New(t)

	q := question{
		one: 4,
		res: []int{2, 2},
	}

	ast.Equal(q.res, constructRectangle(q.one))
}
