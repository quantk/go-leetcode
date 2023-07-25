package main

func containsDuplicate(nums []int) bool {
	m := map[int]bool{}

	for _, val := range nums {
		if _, ok := m[val]; ok {
			return true
		}

		m[val] = true
	}

	return false
}
