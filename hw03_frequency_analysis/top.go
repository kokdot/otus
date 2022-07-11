package main

import (
	// "fmt"
	"strings"
	"sort"
)
const ConstNumbers = 10


func Top10(s string) []string {
	sliceResult := make([]string, 0)
	if s == "" {
		return sliceResult
	}
	sliseForString := strings.Fields(s) //[one, two, three]
	// fmt.Println(sliseForString)
	mapFrequency := make(map[string]int) //[one: 2, two: 3, three:2]
	for _, word := range sliseForString {
		if _, ok := mapFrequency[word]; ok {
			mapFrequency[word]++
		} else {
			mapFrequency[word] = 1
		}
	}
	mapNumbers := make(map[int][]string, 0) // [2:[one, two], 3:[three, for]]
	for key, value := range mapFrequency {
		if _, ok := mapNumbers[value]; ok {
			mapNumbers[value] = append(mapNumbers[value], key)
		} else {
		mapNumbers[value] = []string{key}
		}
	}
	sliceNumbers := make([]int, 0) //[2, 3]
	for key, value := range mapNumbers {
		sort.Strings(value)
		sliceNumbers = append(sliceNumbers, key)
	}
	sort.Ints(sliceNumbers)
	// fmt.Println("sliceNumbers", sliceNumbers)
	Numbers := 0
	for i := 0; i < ConstNumbers; i++ {
		// fmt.Println(i)
		key:= sliceNumbers[len(sliceNumbers) - 1 - i]
		// fmt.Println("key", key)
		// fmt.Println("mapNumbers" ,mapNumbers)
		for _, v := range mapNumbers[key] {
			sliceResult = append(sliceResult, v)
			// fmt.Println("sliseResult", sliceResult)
			Numbers++
			if Numbers == ConstNumbers {
				return sliceResult
			}
		}
		// key++
	}
	return sliceResult
}