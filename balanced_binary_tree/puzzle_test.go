package main

import "testing"

func TestIsBalanced(t *testing.T) {
	input := &TreeNode{
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
	}

	have := isBalanced(input)

	if true != have {
		t.Errorf("Expected %t got %t", true, have)
	}
}
