package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	p1 := head

	if head.Next == nil || head.Next.Next == nil {
		return false
	}
	p2 := head.Next.Next

	for p1.Next != nil && p2 != nil && p2.Next != nil {
		if p1 == p2 {
			return true
		}

		p1 = p1.Next
		p2 = p2.Next.Next
	}

	return false
}
