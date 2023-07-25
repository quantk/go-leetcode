package main

import "testing"

func TestLongestPalindrome(t *testing.T) {
	if longestPalindrome("abccccdd") != 7 {
		t.Errorf("Failed")
	}
}
