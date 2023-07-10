package main

/**
要点：将链表分为两部分，一部分小于x，一部分大于等于x，然后将两部分拼接起来
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}
	var left, right *ListNode
	var l, r *ListNode
	for head != nil {
		if head.Val < x {
			if left == nil {
				left = head
				l = left
			} else {
				left.Next = head
				left = left.Next
			}
		} else {
			if right == nil {
				right = head
				r = right
			} else {
				right.Next = head
				right = right.Next
			}
		}
		head = head.Next
	}
	if left == nil {
		return r
	}
	left.Next = r
	if right != nil {
		right.Next = nil
	}
	return l
}
