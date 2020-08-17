package _002AddTwoSum

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

type para struct {
	one *ListNode
	two *ListNode
}

type ans struct {
	one *ListNode
}

type question struct {
	p para
	a ans
}

func genListNode(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	res := &ListNode{
		Val: nums[0],
	}
	temp := res

	for i := 1; i < len(nums); i++ {
		temp.Next = &ListNode{Val: nums[i]}
		temp = temp.Next
	}
	return res
}

func TestAddTwoNumbers(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		{
			p: para {
				one: genListNode([]int{2, 3, 4}),
				two: genListNode([]int{5, 6, 4}),
			},
			a: ans {
				one: genListNode([]int{7, 9, 8}),
			},
		},
		{
			p: para {
				one: genListNode([]int{9, 8, 7, 6, 5}),
				two: genListNode([]int{1, 1, 2, 3, 4}),
			},
			a: ans {
				one: genListNode([]int{0, 0, 0, 0, 0, 1}),
			},
		},
		{
			p: para {
				one: genListNode([]int{0}),
				two: genListNode([]int{5, 6, 4}),
			},
			a: ans {
				one: genListNode([]int{5, 6, 4}),
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, addTwoNumbers(p.one, p.two), "输入: %v", p)
	}
}