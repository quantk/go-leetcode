package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Stack[T TreeNode] []*T

func (s *Stack[T]) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack[T]) Push(val *T) {
	*s = append(*s, val)
}

func (s *Stack[T]) Pop() (*T, error) {
	if s.IsEmpty() {
		var zero *T
		return zero, fmt.Errorf("empty")
	} else {
		index := len(*s) - 1
		value := (*s)[index]
		*s = (*s)[:index]
		return value, nil
	}
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	s := Stack[TreeNode]{}
	node := root
	s.Push(node)

	for s.IsEmpty() == false {
		node, _ := s.Pop()

		tmp := node.Left
		node.Left = node.Right
		node.Right = tmp

		if node.Left != nil {
			s.Push(node.Left)
		}
		if node.Right != nil {
			s.Push(node.Right)
		}
	}

	return root
}

func invertTreeRec(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	tmp := root.Left
	root.Left = root.Right
	root.Right = tmp

	invertTreeRec(root.Left)
	invertTreeRec(root.Right)

	return root
}
