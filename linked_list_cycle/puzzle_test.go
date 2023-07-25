package main

import "testing"

func TestHasCycle(t *testing.T) {
	p2 := &ListNode{
		Val: 2,
	}
	p4 := &ListNode{
		Val: -4,
	}
	p0 := &ListNode{
		Val: 0,
	}

	p2.Next = p0
	p0.Next = p4
	p4.Next = p2

	input := &ListNode{
		Val:  3,
		Next: p2,
	}

	if hasCycle(input) != true {
		t.Errorf("Failed")
	}
}
