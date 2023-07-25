package main

import "testing"

func TestSearch(t *testing.T) {
	if search([]int{5}, 5) != 0 {
		t.Error("failed")
	}
	if search([]int{-1, 0, 3, 5, 9, 12}, 9) != 4 {
		t.Error("failed")
	}
	if search([]int{-1, 0, 3, 5, 9, 12}, 2) != -1 {
		t.Error("failed")
	}
}
