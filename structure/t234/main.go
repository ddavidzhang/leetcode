package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
* 要点：利用递归达到从后往前遍历的效果，但有个缺点就是过了中点仍然在递归，浪费了时间
 */
func isPalindrome(head *ListNode) bool {
	p := head
	var t func(h *ListNode) bool
	t = func(h *ListNode) bool {
		if h == nil {
			return true
		}
		b := t(h.Next)
		if !b || h.Val != p.Val {
			return false
		}
		p = p.Next
		return true
	}
	return t(head)
}

func main() {
	// 1->2->2->1
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next = &ListNode{Val: 1}
	println(isPalindrome(head))
}

func reverseList(head *ListNode) *ListNode {
	var prev, cur *ListNode = nil, head
	for cur != nil {
		nextTmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = nextTmp
	}
	return prev
}

func endOfFirstHalf(head *ListNode) *ListNode {
	fast := head
	slow := head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

// 快慢指针找中点，然后反转后部分链表，再比较
func isPalindrome2(head *ListNode) bool {
	if head == nil {
		return true
	}

	// 找到前半部分链表的尾节点并反转后半部分链表
	firstHalfEnd := endOfFirstHalf(head)
	secondHalfStart := reverseList(firstHalfEnd.Next)

	// 判断是否回文
	p1 := head
	p2 := secondHalfStart
	result := true
	for result && p2 != nil { // 细节：如果是奇数个节点，第二个链表会比第一个链表少一个节点
		if p1.Val != p2.Val {
			result = false
		}
		p1 = p1.Next
		p2 = p2.Next
	}

	// 还原链表并返回结果
	firstHalfEnd.Next = reverseList(secondHalfStart)
	return result
}
