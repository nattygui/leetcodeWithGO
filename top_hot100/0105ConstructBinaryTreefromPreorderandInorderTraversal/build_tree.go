package buildtree

// TreeNode 二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := newNode(preorder[0])
	rootIndex := findIndex(inorder, preorder[0])
	if rootIndex == -1 {
		return nil
	}
	leftPreorder := preorder[1 : rootIndex+1]
	rightPreorder := preorder[rootIndex+1:]

	leftInorder := inorder[:rootIndex]
	rightInorder := inorder[rootIndex+1:]

	leftTree := buildTree(leftPreorder, leftInorder)
	rightTree := buildTree(rightPreorder, rightInorder)

	root.Left = leftTree
	root.Right = rightTree
	return root
}

func newNode(val int) *TreeNode {
	return &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}

func findIndex(inorder []int, val int) int {
	for index, num := range inorder {
		if val == num {
			return index
		}
	}
	return -1
}
