package main

import (
	"fmt"
	"reflect"
	"testing"
)

func listNodeToArr(node ListNode) []int {
	var result []int
	for node.Next != nil {
		result = append(result, node.Val)
		node = *node.Next
	}

	result = append(result, node.Val)

	return result
}

func arrToListNode(arr []int) *ListNode {
	var root ListNode
	node := &root
	count := len(arr)
	for i, val := range arr {
		node.Val = val
		if i == count-1 {
			break
		}

		nextNode := &ListNode{
			Val:  0,
			Next: nil,
		}
		node.Next = nextNode
		node = nextNode
	}

	return &root
}

func TestListNodeToArr(t *testing.T) {
	arr := listNodeToArr(ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	})
	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("Failed")
	}
}

func TestArrToListNode(t *testing.T) {
	node := *arrToListNode([]int{1, 2, 3})
	fmt.Printf("Node: %v\n", node)
	result := listNodeToArr(node)
	fmt.Printf("Result: %v\n", result)
	if !reflect.DeepEqual(
		result,
		[]int{1, 2, 3},
	) {
		t.Errorf("Failed")
	}
}

func TestMergeTwoLists(t *testing.T) {
	if !reflect.DeepEqual(
		listNodeToArr(*mergeTwoLists(
			arrToListNode([]int{1, 2, 4}),
			arrToListNode([]int{1, 3, 4})),
		),
		[]int{1, 1, 2, 3, 4, 4},
	) {
		t.Error("Failed")
	}

	if !reflect.DeepEqual(
		listNodeToArr(*mergeTwoLists(
			arrToListNode([]int{-9, 3}),
			arrToListNode([]int{5, 7})),
		),
		[]int{-9, 3, 5, 7},
	) {
		t.Error("Failed")
	}
}
