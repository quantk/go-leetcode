package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}
	var current = head
	var prevNode *ListNode

	for current != nil && current.Next != nil {
		next := current.Next
		nextNext := next.Next
		next.Next = current
		current.Next = prevNode
		prevNode = next
		current = nextNext
	}

	if current != nil {
		current.Next = prevNode
		prevNode = current
	}

	return prevNode
}

func reverseListRec(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	last := reverseListRec(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}
