package main

import (
	"reflect"
	"testing"
)

func TestFloodFill1(t *testing.T) {
	input := [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
	expected := [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}}
	if reflect.DeepEqual(floodFill(input, 1, 1, 2), expected) {
		t.Error("Failed")
	}
}
