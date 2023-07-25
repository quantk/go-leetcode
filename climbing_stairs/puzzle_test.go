package main

import "testing"

func TestClimbStairs(t *testing.T) {
	if climbStairs(2) != 2 {
		t.Errorf("failed")
	}
	if climbStairs(5) != 8 {
		t.Errorf("failed")
	}
	if climbStairs(6) != 13 {
		t.Errorf("failed")
	}
	if climbStairs(7) != 21 {
		t.Errorf("failed")
	}
}
