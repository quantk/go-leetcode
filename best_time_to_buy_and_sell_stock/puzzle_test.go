package main

import "testing"

func TestMaxProfit(t *testing.T) {
	if 5 != maxProfit([]int{7, 1, 5, 3, 6, 4}) {
		t.Errorf("failed")
	}

	if 0 != maxProfit([]int{7, 6, 4, 3, 1}) {
		t.Errorf("failed")
	}
}
