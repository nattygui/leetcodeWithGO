package _653TwoSumIVInputisaBST

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	cacheValueMap := make(map[int]bool)
	return dfs(root, k, cacheValueMap)
}

func dfs(root *TreeNode, k int, cache map[int]bool) bool {
	if root == nil {
		return false
	}
	if cache[k - root.Val] {
		return true
	}
	cache[root.Val] = true
	return dfs(root.Left, k, cache) || dfs(root.Right, k, cache)
}