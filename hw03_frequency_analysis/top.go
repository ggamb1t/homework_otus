package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(str string) []string {
	frequencyMap := make(map[string]int64)
	result := make([]string, 0, 10)

	words := strings.Split(str, " ")

	for _, word := range words {
		if word == " " || word == "" {
			continue
		}
		if _, ok := frequencyMap[word]; !ok {
			frequencyMap[word] = 1
		} else {
			frequencyMap[word]++
		}
	}
	type mapToSlice struct {
		Key   string
		Value int64
	}
	var values []mapToSlice
	for key, value := range frequencyMap {
		values = append(values, mapToSlice{key, value})
	}

	sort.Slice(values, func(i, j int) bool {
		return values[i].Value > values[j].Value
	})

	for i := range values {
		if i < 10 {
			result = append(result, values[i].Key)
		} else {
			break
		}
	}

	return result
}
