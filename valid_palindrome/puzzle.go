package main

import (
	"strings"
)

func isPalindrome(s string) bool {
	cleanStringBuilder := strings.Builder{}

	for _, char := range s {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || (char >= '0' && char <= '9') {
			cleanStringBuilder.WriteString(strings.ToLower(string(char)))
		}
	}

	cleanString := cleanStringBuilder.String()

	lp := 0
	rp := len(cleanString) - 1

	for lp <= rp {
		if cleanString[lp] != cleanString[rp] {
			return false
		}

		lp++
		rp--
	}

	return true
}
