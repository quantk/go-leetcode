package main

func min(vals ...int) int {
	min := vals[0]
	for _, val := range vals {
		if val < min {
			min = val
		}
	}

	return min
}

func max(vals ...int) int {
	max := vals[0]
	for _, val := range vals {
		if val > max {
			max = val
		}
	}

	return max
}

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	minPrice := prices[0]
	maxProfit := 0

	for _, val := range prices {
		minPrice = min(val, minPrice)
		maxProfit = max(maxProfit, val-minPrice)
	}

	return maxProfit
}
