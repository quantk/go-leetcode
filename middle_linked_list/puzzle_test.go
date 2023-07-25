package main

import "testing"

func TestMiddleNode(t *testing.T) {
	thirdNode := &ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 5,
			},
		},
	}
	if middleNode(&ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: thirdNode,
		},
	}) != thirdNode {
		t.Errorf("failed")
	}
}
