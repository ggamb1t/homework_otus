package hw03frequencyanalysis

import (
	"fmt"
	"sort"
	"strings"
	"unicode/utf8"
)

func Top10(str string) []string {
	if len(str) == 0 {
		return nil
	}

	var frequencyMap = make(map[string]int64, utf8.RuneCountInString(str))

	// split by [' ', '\n', '\t'] separators
	words := strings.FieldsFunc(str, split)

	// filling map with words as keys and counting them
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

	values := make([]mapToSlice, 0, len(frequencyMap))
	values = append(values, mapToSliceString(frequencyMap)...)
	sort.Slice(values, func(i, j int) bool { // sort by frequency
		return values[i].Count > values[j].Count
	})

	substringsToSort := splitToSubstrings(values)

	result := substringsLexSort(substringsToSort)

	return getFirstNKeys(result, 10)
}

type mapToSlice struct {
	Key   string
	Count int64
}

func getFirstNKeys(values []mapToSlice, n int) []string {
	var result []string
	for i := 0; i < n; i++ {
		result = append(result, values[i].Key)
	}
	return result
}

func mapToSliceString(input map[string]int64) []mapToSlice {
	var values []mapToSlice
	for key, value := range input {
		values = append(values, mapToSlice{
			Key:   key,
			Count: value,
		})
	}
	return values
}

func splitToSubstrings(values []mapToSlice) [][]mapToSlice {
	var (
		prevCount        = values[0].Count
		substringIndex   = 0
		substringsToSort = make([][]mapToSlice, 0)
	)

	substringsToSort = append(substringsToSort, []mapToSlice{})
	for _, value := range values {
		if prevCount != value.Count {
			substringsToSort = append(substringsToSort, []mapToSlice{})
			substringIndex++
			prevCount = value.Count
		}
		substringsToSort[substringIndex] = append(substringsToSort[substringIndex], value)
	}
	return substringsToSort
}

func substringsLexSort(substrings [][]mapToSlice) []mapToSlice {
	var result []mapToSlice
	for _, strs := range substrings {
		for _, str := range strs {
			fmt.Printf("%s - %d\n", str.Key, str.Count)
		}
		sort.Slice(strs, func(i, j int) bool {
			return strs[i].Key < strs[j].Key
		})
		result = append(result, strs...)
	}
	return result
}

func split(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n'
}
