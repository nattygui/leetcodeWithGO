package _021MergeTwoSortedLists

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type param struct {
	one *ListNode
	two *ListNode
}

type question struct {
	params param
	result *ListNode
}

func Test_ok(t *testing.T) {
	ast := assert.New(t)

	questions := []question{
		{
			params: param{
				one: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 4,
							Next: nil,
						},
					},
				},
				two: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: nil,
						},
					},
				},
			},
			result: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 4,
									Next: nil,
								},
							},
						},
					},
				},
			},
		},
	}

	for _, q := range questions {
		ast.Equal(q.result, mergeTwoLists(q.params.one, q.params.two), "输入为：%v， %v", q.params.one, q.params.two)
	}
}