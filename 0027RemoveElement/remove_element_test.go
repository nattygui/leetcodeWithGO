package _027RemoveElement

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type param struct {
	one []int
	two int
}

type question struct {
	params param
	res int
}

func TestOK(t *testing.T) {
	ast := assert.New(t)

	questions := []question{
		{
			params: param{
				one: []int{3,2,2,3},
				two: 3,
			},
			res: 2,
		},
		{
			params: param{
				one: []int{0, 1, 2, 2, 3, 0, 4, 2},
				two: 2,
			},
			res: 5,
		},
	}
	for _, q := range questions {
		ast.Equal(q.res, removeElement(q.params.one, q.params.two), "输入为：%v", q.params)
	}
}