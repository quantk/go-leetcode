package main

import "testing"

func TestCanConstruct(t *testing.T) {
	if canConstruct("a", "b") != false {
		t.Errorf("failed")
	}
	if canConstruct("aa", "ab") != false {
		t.Errorf("failed")
	}
	if canConstruct("aa", "aab") != true {
		t.Errorf("failed")
	}
}
