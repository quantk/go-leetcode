package main

import "testing"

func TestMaxDepth(t *testing.T) {
	if maxDepth(&TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}) != 3 {
		t.Errorf("failed")
	}
}
