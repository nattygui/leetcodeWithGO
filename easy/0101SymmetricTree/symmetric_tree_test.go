package _101SymmetricTree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type question struct {
	one *TreeNode
	ans bool
}


func TestOK(t *testing.T) {
	ast := assert.New(t)

	questions := []question{
		{
			one: &TreeNode{
				Val:   1,
				Left:  &TreeNode{
					Val:   2,
					Left:  &TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   4,
						Left:  nil,
						Right: nil,
					},
				},
				Right: &TreeNode{
					Val:   2,
					Left:  &TreeNode{
						Val:   4,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil,
					},
				},
			},
			ans: true,
		},
		{
			one: &TreeNode{
				Val:   1,
				Left:  &TreeNode{
					Val:   2,
					Left:  nil,
					Right: &TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil,
					},
				},
				Right: &TreeNode{
					Val:   2,
					Left:  nil,
					Right: &TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil,
					},
				},
			},
			ans: false,
		},
	}

	for _, q := range questions {
		ast.Equal(q.ans, isSymmetric(q.one), "输入为%v", q.one)
	}
}