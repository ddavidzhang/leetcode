package t222

/**
要点：其左右子树总有一颗是满二叉树，可以简化计算量
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	l, r := root, root
	hl, hr := 0, 0
	for l != nil {
		l = l.Left
		hl++
	}
	for r != nil {
		r = r.Right
		hr++
	}
	if hl == hr {
		return 1<<hl - 1
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}
