package _530MinimumAbsoluteDifferenceInBST

import "math"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

var min int
var values []int

func getMinimumDifference(root *TreeNode) int {
	min, values = math.MaxInt64, []int{}
	dfs(root)

	for i := 1; i <= len(values); i++ {
		if values[i] - values[i-1] < min {
			min = values[i] - values[i-1]
		}
	}

	return min
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	dfs(root.Left)
	values = append(values, root.Val)
	dfs(root.Right)
}
