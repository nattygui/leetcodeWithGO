# 105. 从前序与中序遍历序列构造二叉树
根据一棵树的前序遍历与中序遍历构造二叉树。

注意:
你可以假设树中没有重复的元素。

例如，给出

```
前序遍历 preorder = [3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]
返回如下的二叉树：

    3
   / \
  9  20
    /  \
   15   7
```

## 解法一

```go
// TreeNode 二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	return buildNode(preorder, inorder)
}

func newNode(val int) *TreeNode {
	return &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}

func buildNode(preorder []int, inorder []int) *TreeNode {
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

	leftTree := buildNode(leftPreorder, leftInorder)
	rightTree := buildNode(rightPreorder, rightInorder)

	root.Left = leftTree
	root.Right = rightTree
	return root
}

func findIndex(preorder []int, val int) int {
	for index, num := range preorder {
		if val == num {
			return index
		}
	}
	return -1
}
```

python
```py
class TreeNode:
    def __init__(self, x):
        self.val = x
        self.left = None
        self.right = None
        
class Solution:
    def buildTree(self, preorder: List[int], inorder: List[int]) -> TreeNode:
        if len(preorder) == 0:
            return None
        root = self.newNode(preorder[0])

        rootIndex = self.findIndex(inorder, preorder[0])

        leftPreorder = preorder[1: rootIndex+1]
        rightPreorder = preorder[rootIndex+1:]

        leftInorder = inorder[:rootIndex]
        rightInorder = inorder[rootIndex+1:]

        leftTree = self.buildTree(leftPreorder, leftInorder)
        rightTree = self.buildTree(rightPreorder, rightInorder)

        root.left = leftTree
        root.right = rightTree

        return root

    
    def newNode(self, val):
        return TreeNode(val)

    
    def findIndex(self, inorder, val):
        for index, num in enumerate(inorder):
            if num == val:
                return index
            
        return -1
```