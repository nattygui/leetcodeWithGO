package symmetric

// TreeNode 二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	leftStack := make([]*TreeNode, 0)
	left := root.Left
	rightStack := make([]*TreeNode, 0)
	right := root.Right
	for {
		for left != nil {
			leftStack = append(leftStack, left)
			left = left.Left
		}
		for right != nil {
			rightStack = append(rightStack, right)
			right = right.Right
		}

		if len(leftStack) == 0 && len(rightStack) == 0 {
			break
		}

		if len(leftStack) != len(rightStack) {
			return false
		}

		left = leftStack[len(leftStack)-1]
		leftStack = leftStack[:len(leftStack)-1]

		right = rightStack[len(rightStack)-1]
		rightStack = rightStack[:len(rightStack)-1]

		if left.Val != right.Val {
			return false
		}

		left = left.Right
		right = right.Left

	}
	return true
}
