package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

const ConstNumbers = 10

func Top10(s string) []string {
	sliceResult := make([]string, 0)
	if s == "" {
		return sliceResult
	}
	sliseForString := strings.Fields(s)  // [one, two, three]
	mapFrequency := make(map[string]int) // [one: 2, two: 3, three:2]
	for _, word := range sliseForString {
		mapFrequency[word]++
		// if _, ok := mapFrequency[word]; ok {
		// 	mapFrequency[word]++
		// } else {
		// 	mapFrequency[word] = 1
		// }
	}
	mapNumbers := make(map[int][]string) // [2:[one, two], 3:[three, for]]
	for key, value := range mapFrequency {
		mapNumbers[value] = append(mapNumbers[value], key)
		// if _, ok := mapNumbers[value]; ok {
		// 	mapNumbers[value] = append(mapNumbers[value], key)
		// } else {
		// 	mapNumbers[value] = []string{key}
		// }
	}
	sliceNumbers := make([]int, 0) // [2, 3]
	for key, value := range mapNumbers {
		sort.Strings(value)
		sliceNumbers = append(sliceNumbers, key)
	}
	sort.Ints(sliceNumbers)
	Numbers := 0
	for i := 0; i < ConstNumbers; i++ {
		key := sliceNumbers[len(sliceNumbers)-1-i]
		for _, v := range mapNumbers[key] {
			sliceResult = append(sliceResult, v)
			Numbers++
			if Numbers == ConstNumbers {
				return sliceResult
			}
		}
	}
	return sliceResult
}
