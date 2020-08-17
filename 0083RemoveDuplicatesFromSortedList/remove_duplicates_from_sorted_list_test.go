package _083RemoveDuplicatesFromSortedList

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type question struct {
	one *ListNode
	res *ListNode
}

func TestOK(t *testing.T) {
	ast := assert.New(t)

	questions := []question{
		{
			one: &ListNode{
				Val:  1,
				Next: &ListNode{
					Val:  1,
					Next: &ListNode{
						Val:  2,
						Next: &ListNode{
							Val:  3,
							Next: &ListNode{
								Val:  3,
								Next: nil,
							},
						},
					},
				},
			},
			res: &ListNode{
				Val:  1,
				Next: &ListNode{
					Val:  2,
					Next: &ListNode{
						Val:  3,
						Next: nil,
					},
				},
			},
		},
	}

	for _, q := range questions {
		ast.Equal(q.res, deleteDuplicates1(q.one), "输入为: %v", q.one)
	}
}
