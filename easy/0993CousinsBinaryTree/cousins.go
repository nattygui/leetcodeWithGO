package iscousins

// TreeNode 二叉树
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCousins(root *TreeNode, x int, y int) bool {
	dmap := make(map[int]int)
	pmap := make(map[int]*TreeNode)
	dfs(root, nil, 0, &dmap, &pmap)
	return dmap[x] == dmap[y] && pmap[x] != pmap[y]
}

func dfs(root, parent *TreeNode, d int, dm *map[int]int, pm *map[int]*TreeNode) {
	if root == nil {
		return
	}
	(*dm)[root.Val] = d
	(*pm)[root.Val] = parent
	dfs(root.Left, root, d+1, dm, pm)
	dfs(root.Right, root, d+1, dm, pm)
}
