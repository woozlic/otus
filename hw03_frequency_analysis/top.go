package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(s string) []string {
	words := strings.Fields(s)
	top := make(map[string]int)
	for _, w := range words {
		if w == "" {
			continue
		}
		top[w]++
	}

	keys := make([]string, 0)
	for key := range top {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		switch {
		case top[keys[i]] > top[keys[j]]:
			return true
		case top[keys[i]] == top[keys[j]]:
			return keys[i] < keys[j]
		default:
			return false
		}
	})
	var topCount int
	if len(top) > 10 {
		topCount = 10
	} else {
		topCount = len(top)
	}
	return keys[:topCount]
}
