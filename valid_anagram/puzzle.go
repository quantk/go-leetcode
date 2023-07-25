package main

func strToMap(s string) map[rune]int {
	result := map[rune]int{}
	for _, c := range s {
		if _, ok := result[c]; !ok {
			result[c] = 1
		} else {
			result[c] = result[c] + 1
		}
	}

	return result
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sMap := strToMap(s)
	tMap := strToMap(t)

	if len(sMap) != len(tMap) {
		return false
	}

	for k, v := range sMap {
		val, ok := tMap[k]
		if !ok || val != v {
			return false
		}
	}

	return true
}
