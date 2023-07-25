package main

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	if reflect.DeepEqual(twoSum([]int{2, 7, 11, 15}, 9), []int{0, 1}) == false {
		t.Error("Failed")
	}

	if reflect.DeepEqual(twoSum([]int{3, 2, 4}, 6), []int{1, 2}) == false {
		t.Error("Failed")
	}

	if reflect.DeepEqual(twoSum([]int{3, 3}, 6), []int{0, 1}) == false {
		t.Error("Failed")
	}
}
