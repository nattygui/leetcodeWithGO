package _014LongestCommonPrefix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type question struct {
	one []string
	res string
}

func Test_ok(t *testing.T) {
	ast := assert.New(t)

	questions := []question{
		{
			one: []string{
				"flower",
				"flow",
				"flight",
			},
			res: "fl",
		},
		{
			one: []string{
				"dog",
				"racecar",
				"car",
			},
			res: "",
		},
	}

	for _, q := range questions {
		ast.Equal(q.res, longestCommonPrefix(q.one), "输入为：%v", q.one)
	}
}