package _257BinaryTreePaths

import "strconv"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	paths := make([]string, 0)
	search(root, "", &paths)
	return paths
}

func search(root *TreeNode, path string, paths *[]string) {
	if root != nil {
		path += strconv.Itoa(root.Val)
		if root.Left == nil && root.Right == nil {
			*paths = append(*paths, path)
		} else {
			path += "->"
			search(root.Left, path, paths)
			search(root.Right, path, paths)
		}
	}
}