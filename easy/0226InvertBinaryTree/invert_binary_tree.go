package _226InvertBinaryTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归 交换所有同一节点下左右子节点
func invertTreeOne(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	invertTreeOne(root.Right)
	invertTreeOne(root.Left)

	temp := root.Left
	root.Left = root.Right
	root.Right = temp

	return root
}

// 迭代
func invertTreeTwo(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		current := queue[0]
		queue = queue[1:len(queue)]
		temp := current.Left
		current.Left = current.Right
		current.Right = temp

		if current.Left != nil {
			queue = append(queue, current.Left)
		}

		if current.Right != nil {
			queue = append(queue, current.Right)

		}
	}

	return root
}
