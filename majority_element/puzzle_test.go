package main

import "testing"

func TestMajorityElement(t *testing.T) {
	if majorityElement([]int{3, 2, 3}) != 3 {
		t.Error("failed")
	}

	if majorityElement([]int{2, 2, 1, 1, 1, 2, 2}) != 2 {
		t.Error("failed")
	}
}
func TestMajorityElementNoSpace(t *testing.T) {
	if majorityElementNoSpace([]int{3, 2, 3}) != 3 {
		t.Error("failed")
	}

	if majorityElementNoSpace([]int{2, 2, 1, 1, 1, 2, 2}) != 2 {
		t.Error("failed")
	}
}
