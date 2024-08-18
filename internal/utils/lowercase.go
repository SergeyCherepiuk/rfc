package utils

import "strings"

func ToLowercaseWords(words []string) []string {
	for i, word := range words {
		words[i] = strings.ToLower(word)
	}
	return words
}
