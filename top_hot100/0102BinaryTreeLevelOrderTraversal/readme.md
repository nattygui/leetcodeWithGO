# 102. 二叉树的层序遍历
给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。

示例：
```
二叉树：[3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层次遍历结果：

[
  [3],
  [9,20],
  [15,7]
]
```

## 解法

层次遍历，构建一个节点切片用来存储当前所有的节点，通过循环获取下一层节点，当节点为空时返回

```go
// TreeNode 二叉树节点
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	if root == nil {
		return result
	}
	curLevelNode := make([]*TreeNode, 0)
	curLevelNode = append(curLevelNode, root)
	for {
		if len(curLevelNode) == 0 {
			break
		}
		nextLevelNodes := make([]*TreeNode, 0)
		curLevelValues := make([]int, 0, len(curLevelNode))
		for _, oneNode := range curLevelNode {
			curLevelValues = append(curLevelValues, oneNode.Val)
			if oneNode.Left != nil {
				nextLevelNodes = append(nextLevelNodes, oneNode.Left)
			}
			if oneNode.Right != nil {
				nextLevelNodes = append(nextLevelNodes, oneNode.Right)
			}
		}
		curLevelNode = nextLevelNodes
		result = append(result, curLevelValues)
	}
	return result
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
    def levelOrder(self, root: TreeNode) -> List[List[int]]:
        result = []
        if root == None:
            return result
        
        cur_level_nodes = [root]
        while True:
            if len(cur_level_nodes) == 0:
                break
            next_level_nodes = []
            cur_level_values = []
            for node in cur_level_nodes:
                cur_level_values.append(node.val)
                if node.left != None:
                    next_level_nodes.append(node.left)
                if node.right != None:
                    next_level_nodes.append(node.right)

            result.append(cur_level_values)
            cur_level_nodes = next_level_nodes
        return result
```