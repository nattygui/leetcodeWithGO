package _404SumOfLeftLeaves

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

var sum int

func sumOfLeftLeaves(root *TreeNode) int {
	sum = 0
	recursion(root)
	return sum
}

func recursion(root *TreeNode) {
	if root == nil {
		return
	}
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		sum += root.Left.Val
	}
	recursion(root.Left)
	recursion(root.Right)
}