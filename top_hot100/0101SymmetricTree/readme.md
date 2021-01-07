# 101. 对称二叉树
给定一个二叉树，检查它是否是镜像对称的。

 

例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

    1
   / \
  2   2
 / \ / \
3  4 4  3
 

但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:

    1
   / \
  2   2
   \   \
   3    3

## 解法一

递归方法，通过对比左子树和右子树的每个节点。

### code

golang
```go
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
	return helper(root.Left, root.Right)
}

func helper(node1, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}

	if node1 == nil || node2 == nil {
		return false
	}

	if node1.Val != node2.Val {
		return false
	}

	return helper(node1.Left, node2.Right) && helper(node1.Right, node2.Left)
}
```

## 解法二

前序遍历左子树的所有节点，后序遍历右子树的所有节点，然后对比每个节点

golang
```go

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
```