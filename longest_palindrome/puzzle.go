package main

func strToMap(s string) map[rune]int {
	res := map[rune]int{}
	for _, c := range s {
		res[c] += 1
	}

	return res
}

func longestPalindrome(s string) int {
	m := strToMap(s)

	res := 0
	for _, val := range m {
		res += val / 2 * 2
		if res%2 == 0 && val%2 == 1 {
			res += 1
		}
	}

	return res
}
