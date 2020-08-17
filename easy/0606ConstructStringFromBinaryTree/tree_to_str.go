package _606ConstructStringFromBinaryTree

import "strconv"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func tree2str(t *TreeNode) string {
	if t == nil {
		return ""
	}
	res := []byte{}
	res = append(res, intToByte(t.Val)...)
	if t.Left != nil {
		preOrder(t.Left, &res)
	} else if t.Right != nil {
		res = append(res, '(')
		res = append(res, ')')
	}
	preOrder(t.Right, &res)
	return string(res)
}

func preOrder(r *TreeNode, path *[]byte) {
	if r == nil {
		return
	}
	*path = append(*path, '(')
	*path = append(*path, intToByte(r.Val)...)
	if r.Left != nil {
		preOrder(r.Left, path)
	} else if r.Right != nil {
		*path = append(*path, '(')
		*path = append(*path, ')')
	}
	preOrder(r.Right, path)
	*path = append(*path, ')')

}

func intToByte(v int) []byte {
	return []byte(strconv.Itoa(v))
}