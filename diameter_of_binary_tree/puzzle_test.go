package main

import "testing"

func TestDiameterOfBinaryTree(t *testing.T) {
	tree := diameterOfBinaryTree(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
	})
	if tree != 1 {
		t.Errorf("failed")
	}
	if diameterOfBinaryTree(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}) != 3 {
		t.Errorf("failed")
	}
}
