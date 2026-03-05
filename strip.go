package main

import (
	"slices"
	"strings"
)

var keywords = []string{
	"kerfuffle",
	"sharbert",
	"fornax",
}

func stripKeywords(body string, keys []string) string {
	words := strings.Split(body, " ")
	values := make([]string, 0, len(words))
	for _, word := range words {
		if slices.Contains(keys, strings.ToLower(word)) {
			values = append(values, "****")
		} else {
			values = append(values, word)
		}
	}

	return strings.Join(values, " ")
}
