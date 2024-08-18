package utils

import "strings"

func Trim(cutset string) func([]string) []string {
	return func(words []string) []string {
		for i, word := range words {
			words[i] = strings.Trim(word, "-")
		}
		return words
	}
}
