package main

import (
	"strings"
)

func wordFrequency(text string) map[string]int {
	words := strings.Fields(text)
	frequency := make(map[string]int)
	for _, word := range words {
		frequency[word]++
	}
	return frequency
}
