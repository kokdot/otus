package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

const ConstNumbers = 10

func Top10(s string) []string {
	var sliceResult []string
	if s == "" {
		return sliceResult
	}
	sliseForString := strings.Fields(s)  // [one, two, three]
	mapFrequency := make(map[string]int) // [one: 2, two: 3, three:2]
	for _, word := range sliseForString {
		mapFrequency[word]++
	}

	for key := range mapFrequency {
		sliceResult = append(sliceResult, key)
	}
	sort.Slice(sliceResult, func(i, j int) bool {
		if mapFrequency[sliceResult[i]] == mapFrequency[sliceResult[j]] {
			return sliceResult[i] < sliceResult[j]
		}
		return mapFrequency[sliceResult[i]] > mapFrequency[sliceResult[j]]
	})

	return sliceResult[:ConstNumbers]
}
