package _671SecondMinimumNodeInaBinaryTree

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Result struct {
	first  int
	second int
}

func (r *Result) swap(value int) {
	if value == r.first || value == r.second {
		return
	}

	if value < r.second && value > r.first {
		r.second = value
	} else if value < r.first {
		r.second = r.first
		r.first = value
	}
}

var res *Result

func findSecondMinimumValue(root *TreeNode) int {
	if root == nil {
		return -1
	}
	res = &Result{
		first:  math.MaxInt64,
		second: math.MaxInt64,
	}

	dfs(root)

	if res.first == res.second || res.first == math.MaxInt64 || res.second == math.MaxInt64 {
		return -1
	}
	return res.second
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	res.swap(root.Val)
	dfs(root.Left)
	dfs(root.Right)
}
