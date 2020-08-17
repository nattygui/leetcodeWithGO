package _125ValidPalindrome

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type question struct {
	one string
	ans bool
}

func TestOk(t *testing.T) {
	ast := assert.New(t)

	questions := []question{
		{
			one: "A man, a plan, a canal: Panama",
			ans: true,
		},
		{
			one: "race a car",
			ans: false,
		},
	}

	for _, q := range questions {
		ast.Equal(q.ans, isPalindrome(q.one), "输入为%v", q.one)
	}

}