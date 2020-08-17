package _013RomanToInteger

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type question struct {
	one string
	ans int
}

func Test_ok(t *testing.T) {
	ast := assert.New(t)

	questions := []question{
		{
			one: "III",
			ans: 2,
		},
		{
			one: "IV",
			ans: 4,
		},
		{
			one: "IX",
			ans: 9,
		},
		{
			one: "LVIII",
			ans: 58,
		},
	}

	for _, q := range questions {
		ast.Equal(q.ans, romanToInt(q.one), "输入为：%v", q.one)
	}
}