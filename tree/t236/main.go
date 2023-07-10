package t236

/**
要点：标准的遍历查找而已
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	ant, _, _ := traverse(root, p, q)
	return ant
}

func traverse(root, p, q *TreeNode) (*TreeNode, bool, bool) {
	if root == nil {
		return nil, false, false
	}
	ant, leftP, leftQ := traverse(root.Left, p, q)
	if ant != nil {
		return ant, true, true
	}
	ant, rightP, rightQ := traverse(root.Right, p, q)
	if ant != nil {
		return ant, true, true
	}
	if (leftP && rightQ) || (leftQ && rightP) {
		return root, true, true
	}
	if root == p {
		if leftQ || rightQ {
			return root, true, true
		}
		return nil, true, false
	}
	if root == q {
		if leftP || rightP {
			return root, true, true
		}
		return nil, false, true
	}
	return nil, leftP || rightP, leftQ || rightQ
}
