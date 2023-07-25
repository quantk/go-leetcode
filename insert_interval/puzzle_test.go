package main

import (
	"reflect"
	"testing"
)

func TestInsert1(t *testing.T) {
	got := insert([][]int{{1, 3}, {6, 9}}, []int{2, 5})
	expected := [][]int{{1, 5}, {6, 9}}
	if reflect.DeepEqual(got, expected) != true {
		t.Errorf("failed: expected: %+v, got: %+v", expected, got)
	}
}
func TestInsert2(t *testing.T) {
	got := insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8})
	expected := [][]int{{1, 2}, {3, 10}, {12, 16}}
	if reflect.DeepEqual(got, expected) != true {
		t.Errorf("failed: expected: %+v, got: %+v", expected, got)
	}
}
