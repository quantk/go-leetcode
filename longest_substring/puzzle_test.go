package main

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	if lengthOfLongestSubstring("bbbbb") != 1 {
		t.Errorf("failed")
	}
	if lengthOfLongestSubstring("au") != 2 {
		t.Errorf("failed")
	}
	if lengthOfLongestSubstring("abcabcbb") != 3 {
		t.Errorf("failed")
	}
	if lengthOfLongestSubstring("pwwkew") != 3 {
		t.Errorf("failed")
	}
}
