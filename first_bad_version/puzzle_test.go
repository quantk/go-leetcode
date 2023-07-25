package main

import "testing"

func TestFirstBadVersion(t *testing.T) {
	if firstBadVersion(2) != 1 {
		t.Errorf("failed")
	}
}
