package rangesum

// TreeNode 二叉树
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	sum := 0
	dfs(root, L, R, &sum)
	return sum
}

func dfs(root *TreeNode, L int, R int, sum *int) {
	if root == nil {
		return
	}
	if root.Val >= L {
		dfs(root.Left, L, R, sum)
	}

	if root.Val <= R {
		dfs(root.Right, L, R, sum)
	}
	if root.Val >= L && root.Val <= R {
		*sum += root.Val
	}
}
