package diameter

// TreeNode 二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	_, _, rootMax := dfs(root)
	return rootMax - 1
}

func dfs(root *TreeNode) (int, int, int) {
	if root == nil {
		return 0, 0, 0
	}
	leftSubLeft, leftSubRight, leftPathMax := dfs(root.Left)
	rightSubLeft, rightSubRight, rightPathMax := dfs(root.Right)
	leftMax := max(leftSubLeft, leftSubRight)
	if root.Left != nil {
		leftMax++
	}
	rightMax := max(rightSubLeft, rightSubRight)
	if root.Right != nil {
		rightMax++
	}
	curPathMax := max(leftMax+rightMax, max(leftPathMax, rightPathMax))

	return leftMax, rightMax, curPathMax

	// leftAll := leftSubLeft + leftSubRight
	// rightAll := rightSubLeft + rightSubRight
	// curMax := max(leftSubLeft, leftSubRight) + max(rightSubLeft, rightSubRight) + 1
	// return max(leftSubLeft, leftSubRight) + 1, max(rightSubLeft, rightSubRight) + 1, max(curMax, max(leftAll, rightAll))
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}
