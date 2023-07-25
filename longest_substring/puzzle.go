package main

func lengthOfLongestSubstring(s string) int {
	start := 0
	maxStr := 0
	m := map[byte]int{}

	for end := 0; end < len(s); end++ {
		rc := s[end]
		if _, ok := m[rc]; ok {
			start = max(start, m[rc]+1)
		}
		m[rc] = end
		maxStr = max(maxStr, end-start+1)
	}

	return maxStr
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
