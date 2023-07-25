package main

import "testing"

func TestAddBinary(t *testing.T) {
	if addBinary("1", "111") != "1000" {
		t.Error("failed")
	}
	if addBinary("0", "0") != "0" {
		t.Error("failed")
	}

	if addBinary("11", "1") != "100" {
		t.Error("failed")
	}

	if addBinary("1010", "1011") != "10101" {
		t.Error("failed")
	}
}
