package main

import "testing"

func TestLowestCommonAncestor(t *testing.T) {
	p := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 0,
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
	}
	q := &TreeNode{
		Val: 8,
		Left: &TreeNode{
			Val: 7,
		},
		Right: &TreeNode{
			Val: 9,
		},
	}
	input := &TreeNode{
		Val:   6,
		Left:  p,
		Right: q,
	}

	expected := 6

	have := lowestCommonAncestor(input, p, q).Val
	if have != expected {
		t.Errorf("expected %d got %d", expected, have)
	}
}
