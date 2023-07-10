package t99

/**
要点：
1、了解二叉搜索树的性质，中序遍历的结果是有序的，所以可以通过中序遍历找到两个错误的节点
2、编码上注意指针是值传递，想修改原始指针的值，需要传递指针的指针
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode) {
	var s, t, p *TreeNode
	traverse(root, &p, &s, &t)
	s.Val, t.Val = t.Val, s.Val
}

func traverse(node *TreeNode, p, s, t **TreeNode) {
	if node == nil {
		return
	}
	traverse(node.Left, p, s, t)
	if *p != nil && node.Val < (*p).Val {
		if *s == nil {
			*s = *p
		}
		*t = node
	}
	*p = node
	traverse(node.Right, p, s, t)
}
