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
		_, ok := top[w]
		if !ok {
			top[w] = 1
		} else {
			top[w]++
		}
	}

	keys := make([]string, len(top))
	for key := range top {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		if top[keys[i]] > top[keys[j]] {
			return true
		} else if top[keys[i]] == top[keys[j]] {
			return keys[i] < keys[j]
		} else {
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
