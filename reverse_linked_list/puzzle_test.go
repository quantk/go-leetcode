package main

import (
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
func TestReverseList(t *testing.T) {
	result := reverseList(arrToListNode([]int{1, 2, 3, 4, 5}))
	if reflect.DeepEqual([]int{5, 4, 3, 2, 1}, listNodeToArr(*result)) != true {
		t.Errorf("failed")
	}

	result2 := reverseList(arrToListNode([]int{1, 2}))
	if reflect.DeepEqual([]int{2, 1}, listNodeToArr(*result2)) != true {
		t.Errorf("failed")
	}
	result3 := reverseList(arrToListNode([]int{1, 2, 3, 4}))
	if reflect.DeepEqual([]int{4, 3, 2, 1}, listNodeToArr(*result3)) != true {
		t.Errorf("failed")
	}
}
func TestReverseListRec(t *testing.T) {
	result := reverseListRec(arrToListNode([]int{1, 2, 3, 4, 5}))
	if reflect.DeepEqual([]int{5, 4, 3, 2, 1}, listNodeToArr(*result)) != true {
		t.Errorf("failed")
	}

	result2 := reverseListRec(arrToListNode([]int{1, 2}))
	if reflect.DeepEqual([]int{2, 1}, listNodeToArr(*result2)) != true {
		t.Errorf("failed")
	}
	result3 := reverseListRec(arrToListNode([]int{1, 2, 3, 4}))
	if reflect.DeepEqual([]int{4, 3, 2, 1}, listNodeToArr(*result3)) != true {
		t.Errorf("failed")
	}
}
