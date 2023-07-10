package main

/**
要点：重点就是熟悉二叉树节点的编号方式，并数量掌握二叉树层次遍历的框架
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 层序遍历
	var q []*TreeNode
	root.Val = 0
	q = append(q, root)
	res := 1
	for len(q) > 0 {
		var tq []*TreeNode
		for _, node := range q {
			if node.Left != nil {
				node.Left.Val = 2 * node.Val
				tq = append(tq, node.Left)
			}
			if node.Right != nil {
				node.Right.Val = 2*node.Val + 1
				tq = append(tq, node.Right)
			}
		}
		if len(tq) == 0 {
			break
		}
		res = max(res, tq[len(tq)-1].Val-tq[0].Val+1)
		q = tq
	}
	return res
}
func max(res int, i int) int {
	if res > i {
		return res
	}
	return i
}
