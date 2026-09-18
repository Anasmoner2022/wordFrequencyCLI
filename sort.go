package main

import (
	"bufio"
	"io"
	"sort"
	"strings"
	"unicode"
)

func cleanWord(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) {
			return unicode.ToLower(r)
		}
		return -1
	}, s)
}

func countWords(r io.Reader) (map[string]int, error) {
	counts := make(map[string]int)
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := cleanWord(scanner.Text())
		if word != "" {
			counts[word]++
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return counts, nil
}

func sortedCount(counts map[string]int) []WordCount {
	result := make([]WordCount, 0, len(counts))
	for word, count := range counts {
		result = append(result, WordCount{Word: word, Count: count})
	}
	sort.Slice(result, func(i, j int) bool {
		if result[i].Count != result[j].Count {
			return result[i].Count > result[j].Count
		}
		return result[i].Word < result[j].Word
	})
	return result
}
