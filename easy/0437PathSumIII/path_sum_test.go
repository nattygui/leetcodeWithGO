package _437PathSumIII

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
				Val:   10,
				Left:  &TreeNode{
					Val:   5,
					Left:  &TreeNode{
						Val:   3,
						Left:  &TreeNode{
							Val:   3,
							Left:  nil,
							Right: nil,
						},
						Right: &TreeNode{
							Val:   -2,
							Left:  nil,
							Right: nil,
						},
					},
					Right: &TreeNode{
						Val:   2,
						Left:  nil,
						Right: &TreeNode{
							Val:   1,
							Left:  nil,
							Right: nil,
						},
					},
				},
				Right: &TreeNode{
					Val:   -3,
					Left:  nil,
					Right: &TreeNode{
						Val:   11,
						Left:  nil,
						Right: nil,
					},
				},
			},
			ans: 3,
		},
	}

	for _, q := range questions {
		ast.Equal(q.ans, pathSumOne(q.one, 8))
	}


}
