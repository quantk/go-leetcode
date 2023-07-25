package main

func isBadVersion(v int) bool {
	return v >= 1
}

func firstBadVersion(n int) int {
	l := 1
	r := n

	for l < r {
		m := l + (r-l)/2
		if isBadVersion(m) {
			r = m
		} else {
			l = m + 1
		}
	}

	return l
}
