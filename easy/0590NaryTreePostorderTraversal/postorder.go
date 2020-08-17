package _590NaryTreePostorderTraversal

type Node struct {
	Val int
	Children []*Node
}

func postorder(root *Node) []int {
	if  root == nil {
		return []int{}
	}
	res := []int{}
	stack := []*Node{root}

	for len(stack) != 0 {
		if stack[len(stack)-1].Children == nil {
			res = append(res, stack[len(stack)-1].Val)
			stack = stack[:len(stack)-1]
		} else {
			temp := len(stack)-1
			for i := len(stack[temp].Children) -1; i >= 0; i-- {
				stack = append(stack, stack[temp].Children[i])
			}
			stack[temp].Children = nil
		}
	}

	return res
}
