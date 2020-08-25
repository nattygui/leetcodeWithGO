package mindiff

import "testing"

func TestOK(t *testing.T) {
	root := &TreeNode{
		Val:  27,
		Left: nil,
		Right: &TreeNode{
			Val:  34,
			Left: nil,
			Right: &TreeNode{
				Val: 58,
				Left: &TreeNode{
					Val: 50,
					Left: &TreeNode{
						Val:   44,
						Left:  nil,
						Right: nil,
					},
				},
			},
		},
	}
	res := minDiffInBST(root)
	if res != 6 {
		t.Error("error")
	}
}
