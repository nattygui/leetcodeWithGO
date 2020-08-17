package _501FindModeInBinarySearchTree

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

var res []int
var cur int
var max int
var count int

func findMode(root *TreeNode) []int {
	res, cur, max, count = []int{}, 0, 1, 0
	dfs(root)
	return res
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	dfs(root.Left)
	if root.Val != cur {
		count = 0
	}
	count++
	cur = root.Val
	if max < count {
		max = count
		res = []int{cur}
	} else if max == count {
		res = append(res, cur)
	}
	dfs(root.Right)

}