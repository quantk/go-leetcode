package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	if isPalindrome("A man, a plan, a canal: Panama") != true {
		t.Errorf("failed")
	}

	if isPalindrome("race a car") != false {
		t.Errorf("failed")
	}

	if isPalindrome(" ") != true {
		t.Errorf("failed")
	}

	if isPalindrome("коок") != true {
		t.Errorf("failed")
	}
}
