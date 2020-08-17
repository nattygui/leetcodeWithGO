package _104MaximumDepthOfBinaryTree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type question struct {
	one *TreeNode
	ans int
}

func TestOK(t *testing.T) {
	ast := assert.New(t)
	
	questions := []question{
		{
			one: &TreeNode{
				Val:   3,
				Left:  &TreeNode{
					Val:   9,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   20,
					Left:  &TreeNode{
						Val:   15,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   7,
						Left:  nil,
						Right: nil,
					},
				},
			},
			ans: 3,
		},
	}

	for _, q := range questions {
		ast.Equal(q.ans, maxDepth(q.one), "输入为%v", q.one)
	}
}
