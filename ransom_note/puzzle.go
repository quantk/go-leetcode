package main

func strToMap(str string) map[rune]int {
	res := map[rune]int{}

	for _, c := range str {
		res[c] = res[c] + 1
	}

	return res
}

func canConstruct(ransomNote string, magazine string) bool {
	m1 := strToMap(ransomNote)

	for _, c := range magazine {
		m1[c] = m1[c] - 1
	}

	for _, val := range m1 {
		if val > 0 {
			return false
		}
	}

	return true
}
