package main

func maxSubArray(nums []int) int {
	maxSum := nums[0]
	sum := 0

	for _, val := range nums {
		if sum+val < val {
			sum = val
		} else {
			sum += val
		}

		maxSum = max(maxSum, sum)
	}

	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
