package _234PalindromeLinkedList

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type question struct {
	one *ListNode
	res bool
}

func TestOK(t *testing.T) {
	ast := assert.New(t)

	questions := []question{
		{
			one: &ListNode{
				Val:  1,
				Next: &ListNode{
					Val:  3,
					Next: &ListNode{
						Val:  5,
						Next: &ListNode{
							Val:  5,
							Next: &ListNode{
								Val:  3,
								Next: &ListNode{
									Val:  1,
									Next: nil,
								},
							},
						},
					},
				},
			},
			res: true,
		},
	}

	for _, q := range questions {
		ast.Equal(q.res, isPalindromeTwo(q.one))
	}

}