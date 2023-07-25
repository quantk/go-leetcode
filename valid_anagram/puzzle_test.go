package main

import "testing"

func TestIsAnagram(t *testing.T) {
	if isAnagram("anagram", "nagaram") != true {
		t.Error("Failed")
	}
	if isAnagram("rat", "car") != false {
		t.Error("Failed")
	}
}
