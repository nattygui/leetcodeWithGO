package _028ImplementStrStr

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type param struct {
	one string
	two string
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
				one: "heololllll",
				two: "ll",
			},
			res: 5,
		},
		{
			params: param{
				one: "a",
				two: "a",
			},
			res: 0,
		},
	}

	for _, q := range questions {
		ast.Equal(q.res, strStr(q.params.one, q.params.two), "输入为:%v", q.params)
	}
}
