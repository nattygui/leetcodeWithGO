package _572SubtreeOfAnotherTree

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

var res bool

func isSubtree(s *TreeNode, t *TreeNode) bool {
	res = false
	dfs(s, t)
	return res
}

func dfs(s *TreeNode, t *TreeNode) {
	if s == nil {
		return
	}
	if s.Val == t.Val {
		res = res || recursion(s, t)
	}
	dfs(s.Left, t)
	dfs(s.Right, t)
}

func recursion(root *TreeNode, t *TreeNode) bool {
	if root == nil && t == nil {
		return true
	}
	if root == nil || t == nil || root.Val != t.Val {
		return false
	}
	return recursion(root.Left, t.Left) && recursion(root.Right, t.Right)
}