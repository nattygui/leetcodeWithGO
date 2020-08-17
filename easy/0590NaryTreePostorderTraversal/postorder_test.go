package _590NaryTreePostorderTraversal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type question struct {
	one *Node
	ans []int
}

func TestOk(t *testing.T) {
	ast := assert.New(t)

	q := question{
		one: &Node{
			Val:      1,
			Children: []*Node{
				&Node{
					Val:      3,
					Children: []*Node{
						{
							Val: 5,
							Children: nil,
						},
						{
							Val: 6,
							Children: nil,
						},
					},
				},
				&Node{
					Val:      2,
					Children: nil,
				},
				&Node{
					Val:      4,
					Children: nil,
				},
			},
		},
		ans: []int{5,6,3,2,4,1},
	}

	ast.Equal(q.ans, postorder(q.one))

}