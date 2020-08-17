package _563BinaryTreeTilt

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func findTilt(root *TreeNode) int {
	res := 0
	getTilt(root, &res)
	return res
}

func getTilt(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}
	left := getTilt(root.Left, res)
	right := getTilt(root.Right, res)

	*res += abs(left, right)
	return root.Val + left + right
}

func abs(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}