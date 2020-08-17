package _009PalindromeNumber

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type question struct {
	one int
	res	bool
}

func Test_ok(t *testing.T) {
	ast := assert.New(t)

	questions := []question{
		{
			one: 121,
			res: true,
		},
		{
			one: -111,
			res: false,
		},
		{
			one: 34212,
			res: false,
		},
	}

	for _, q := range questions {

		ast.Equal(q.res, palindromeNumber(q.one), "输入：%v", q.one)
	}
}
