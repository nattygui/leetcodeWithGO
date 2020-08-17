package _589NaryTreePreorderTraversal

type Node struct {
    Val int
    Children []*Node
}

func preorder(root *Node) []int {
	res := []int{}
	stack := []*Node{root}

	for len(stack) != 0 {
		res = append(res, stack[0].Val)
		stack = append(stack[0].Children, stack[1:]...)
	}
	return res
}