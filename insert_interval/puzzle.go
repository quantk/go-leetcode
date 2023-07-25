package main

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func insert(intervals [][]int, newinterval []int) [][]int {
	var res [][]int
	i := 0

	for ; i < len(intervals) && intervals[i][1] < newinterval[0]; i++ {
		res = append(res, intervals[i])
	}

	for ; i < len(intervals) && intervals[i][0] <= newinterval[1]; i++ {
		newinterval[0] = min(newinterval[0], intervals[i][0])
		newinterval[1] = max(newinterval[1], intervals[i][1])
	}

	res = append(res, newinterval)

	for ; i < len(intervals); i++ {
		res = append(res, intervals[i])
	}

	return res
}
