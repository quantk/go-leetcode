package main

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	n1 := 1
	n2 := 1
	result := 0

	for i := 3; i <= n; i++ {
		result = n1 + n2
		n1 = n2
		n2 = result
	}

	return n1 + n2
}
