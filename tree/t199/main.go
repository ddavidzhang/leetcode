package t199

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	var ans []int
	if root == nil {
		return ans
	}
	// 层序遍历
	var q []*TreeNode
	q = append(q, root)
	ans = append(ans, root.Val)
	for len(q) > 0 {
		var tq []*TreeNode
		for _, node := range q {
			if node.Left != nil {
				tq = append(tq, node.Left)
			}
			if node.Right != nil {
				tq = append(tq, node.Right)
			}
		}
		if len(tq) == 0 {
			break
		}
		ans = append(ans, tq[len(tq)-1].Val)
		q = tq
	}
	return ans
}
