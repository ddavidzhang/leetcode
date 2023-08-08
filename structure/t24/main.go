package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	var pre *ListNode
	var cur *ListNode
	var next *ListNode
	if head == nil || head.Next == nil {
		return head
	}
	pre = head
	cur = head.Next
	next = cur.Next
	cur.Next = pre
	pre.Next = swapPairs(next)
	return cur
}
