package dynamicprogramming

// TreeNode 二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}

func min(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
