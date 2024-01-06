package main

import (
	"cmp"
	"slices"
	"strings"
)

func Chunk(data, delimiter string, n int) [][]string {
	cleanLines := strings.ReplaceAll(data, "\r", "")
	splitData := strings.Split(strings.Trim(cleanLines, delimiter), delimiter)

	for i := range splitData {
		splitData[i] = strings.TrimSpace(splitData[i])
	}
	slices.SortFunc(splitData, func(a, b string) int {
		return cmp.Compare(a, b)
	})
	nChunks := len(splitData) / n
	remaining := len(splitData) % n
	if remaining != 0 {
		nChunks++
	}
	chunks := make([][]string, nChunks)
	for i, r := range splitData {
		groupN := i / n
		chunks[groupN] = append(chunks[groupN], r)
	}

	return chunks

}
