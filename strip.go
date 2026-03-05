package main

import (
	"strings"
)

var badKeyWords = map[string]struct{}{
	"kerfuffle": {},
	"sharbert":  {},
	"fornax":    {},
}

func stripKeywords(body string, keys map[string]struct{}) string {
	words := strings.Split(body, " ")
	for i, word := range words {
		if _, ok := keys[strings.ToLower(word)]; ok {
			words[i] = "****"
		}
	}

	return strings.Join(words, " ")
}
