package peterpiper

import (
	"strings"
)

// KMPStringSearch returns a list of starting indexs for a needle found in a haystack
func KMPStringSearch(haystack string, needle string) []int {
	ret := []int{}
	m, n := len(needle), len(haystack)
	if m == 0 || n == 0 || n < m {
		return ret
	}

	haystack = strings.ToLower(haystack)
	needle = strings.ToLower(needle)

	kmpTable := preProcess(haystack)
	i, j := 0, 0
	x, y := []byte(needle), []byte(haystack)

	for j < n {
		for i > -1 && x[i] != y[j] {
			i = kmpTable[i]
		}
		i++
		j++
		if i >= m {
			ret = append(ret, j-i)
			i = kmpTable[i]
		}
	}

	return ret
}

func preProcess(text string) []int {
	i, j, length := 0, -1, len(text)-1
	kmpTable := make([]int, len(text))
	kmpTable[0] = -1

	for i < length {
		for j > -1 && text[i] != text[j] {
			j = kmpTable[j]
		}

		i++
		j++

		if text[i] == text[j] {
			kmpTable[i] = kmpTable[j]
		} else {
			kmpTable[i] = j
		}
	}

	return kmpTable
}
