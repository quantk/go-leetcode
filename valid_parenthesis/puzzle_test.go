package main

import "testing"

func TestIsValid(t *testing.T) {
	if isValid("()") != true {
		t.Error("Failed")
	}

	if isValid("()[]{}") != true {
		t.Error("Failed")
	}

	if isValid("(]") != false {
		t.Error("Failed")
	}

    if isValid("((") != false {
		t.Error("Failed")
	}
}
