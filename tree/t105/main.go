package t105

/**
要点：树的遍历， 中序的遍历确定范围，前序的遍历确定根节点
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	inMap := make(map[int]int)
	for i, v := range inorder {
		inMap[v] = i
	}
	return buildTreeNode(preorder, inorder, 0, len(preorder)-1, 0, len(inorder)-1, inMap)
}

func buildTreeNode(preorder, inorder []int, preStart, preEnd, inStart, inEnd int, inMap map[int]int) *TreeNode {
	if preStart > preEnd || inStart > inEnd {
		return nil
	}
	root := &TreeNode{
		Val: preorder[preStart],
	}
	inRoot := inMap[root.Val]
	numsLeft := inRoot - inStart
	root.Left = buildTreeNode(preorder, inorder, preStart+1, preStart+numsLeft, inStart, inRoot-1, inMap)
	root.Right = buildTreeNode(preorder, inorder, preStart+numsLeft+1, preEnd, inRoot+1, inEnd, inMap)
	return root
}
