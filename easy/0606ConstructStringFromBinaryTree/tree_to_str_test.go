package _606ConstructStringFromBinaryTree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type question struct {
	one *TreeNode
	ans string
}

func TestOK(t *testing.T) {
	ast := assert.New(t)

	q := question{
		one: &TreeNode{
			Val:   1,
			Left:  &TreeNode{
				Val:   2,
				Left:  &TreeNode{
					Val:   4,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
		ans: "1(2(4))(3)",
	}

	ast.Equal(q.ans, tree2str(q.one))

}