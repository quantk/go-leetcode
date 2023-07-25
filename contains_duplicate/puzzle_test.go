package main

import "testing"

func TestContainsDuplicate(t *testing.T) {
	if containsDuplicate([]int{1, 2, 3, 1}) != true {
		t.Errorf("failed")
	}

	if containsDuplicate([]int{1, 2, 3, 4}) != false {
		t.Errorf("failed")
	}

	if containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}) != true {
		t.Errorf("failed")
	}
}
