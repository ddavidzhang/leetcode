package main

import (
	"fmt"
)

/**
  最大陆路径和，拆为"左+右+根"节点
*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	var ant = -300000000
	onSideMax(root, &ant)
	return ant
}

func onSideMax(node *TreeNode, ant *int) int {
	if node == nil {
		return 0
	}
	l := onSideMax(node.Left, ant)
	r := onSideMax(node.Right, ant)
	*ant = max([]int{*ant, l + r + node.Val, l + node.Val, r + node.Val, node.Val})
	return max([]int{l + node.Val, r + node.Val, node.Val})
}

// 求最大
func max(a []int) int {
	m := a[0]
	for _, v := range a[1:] {
		if v > m {
			m = v
		}
	}
	return m
}
func main() {
	n := &TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	}
	sum := maxPathSum(n)
	fmt.Println(sum)
}
