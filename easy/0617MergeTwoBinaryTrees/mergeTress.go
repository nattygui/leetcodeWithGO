package _617MergeTwoBinaryTrees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	res := t1
	dfs(t1, t2)
	return res
}

func dfs(t1 *TreeNode, t2 *TreeNode) {
	if t1 == nil && t2 == nil {
		return
	} else if t1 != nil && t2 != nil {
		t1.Val += t2.Val
	} else if t1 == nil {
		temp := &TreeNode{}
		temp.Val = t2.Val
		t1 = temp
	}
	dfs(t1.Left, t2.Left)
	dfs(t1.Right, t2.Right)
}

func mergeTreesOne(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	t1.Val += t2.Val
	t1.Left = mergeTrees(t1.Left, t2.Left)
	t1.Right = mergeTrees(t1.Right, t2.Right)
	return t1
}
