package main

func search(nums []int, target int) int {
	left := 0
	right := len(nums)

	for left != right {
		middle := (left + right) / 2
		if nums[middle] == target {
			return middle
		}

		if nums[middle] < target {
			left = middle + 1
		} else {
			right = middle
		}
	}
	return -1
}
