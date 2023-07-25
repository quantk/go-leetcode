package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if abs(left-right) > 1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}

func maxDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}

	left := maxDepth(node.Left)
	right := maxDepth(node.Right)

	return max(left, right) + 1
}

func abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
