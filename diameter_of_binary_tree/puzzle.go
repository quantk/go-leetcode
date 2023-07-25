package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var maxDiameter = 0

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	return max(maxDiameter, left+right)
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	maxDiameter = max(maxDiameter, left+right)

	return max(left, right) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
