package _538ConvertBSTToGreaterTree

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

var add int

func convertBST(root *TreeNode) *TreeNode {
	add = 0
	backRecursion(root)
	return root
}

func backRecursion(root *TreeNode) {
	if root == nil {
		return
	}
	backRecursion(root.Left)
	backRecursion(root.Right)
	root.Val += add
	add += root.Val
}
