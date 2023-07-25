package main

import "sort"

func majorityElement(nums []int) int {
	count := len(nums)
	m := map[int]int{}

	for _, val := range nums {
		m[val] += 1
		if m[val] > count/2 {
			return val
		}
	}

	return 0
}

func majorityElementNoSpace(nums []int) int {
	sort.Sort(sort.IntSlice(nums))
	return nums[len(nums)/2]
}
