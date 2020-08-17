package _687LongestUnivaluePath

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res int

func longestUnivaluePath(root *TreeNode) int {
	res = 0
	dfs(root)
	return res
}

func dfs(node *TreeNode) int {
	if node == nil {
		return 0
	}
	l := dfs(node.Left)
	r := dfs(node.Right)
	left, right := 0, 0
	if node.Left != nil && node.Left.Val == node.Val {
		left = l + 1
	}
	if node.Right != nil && node.Right.Val == node.Val {
		right = r + 1
	}
	res = max(res, left+right)
	return max(left, right)
}

func max(x, t int) int {
	if x > t {
		return x
	}
	return t
}
