package mindiff

import "math"

// TreeNode Definition for a binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDiffInBST(root *TreeNode) int {
	var pre *int
	res := math.MaxInt64

	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		dfs(root.Left)
		if pre != nil {
			temp := root.Val - *pre
			if res > temp {
				res = temp
			}
		}
		pre = &root.Val
		dfs(root.Right)
	}
	dfs(root)
	return res
}
