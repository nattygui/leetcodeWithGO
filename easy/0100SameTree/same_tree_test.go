package _100SameTree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type param struct {
	one *TreeNode
	two *TreeNode
}

type question struct {
	params param
	res bool
}

func TestOK(t *testing.T) {
	ast := assert.New(t)

	questions := []question{
		{
			params: param{
				one: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:   2,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   3,
						Left:  nil,
						Right: nil,
					},
				},
				two: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:   2,
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
			res: true,
		},
		{
			params: param{
				one: &TreeNode{
					Val:   1,
					Left:  &TreeNode{
						Val:   2,
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				},
				two: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: &TreeNode{
						Val:   2,
						Left:  nil,
						Right: nil,
					},
				},
			},
			res: false,
		},
	}

	for _, q := range questions {
		ast.Equal(q.res, isSameTree(q.params.one, q.params.two), "输入为: %v", q.params)
	}
}