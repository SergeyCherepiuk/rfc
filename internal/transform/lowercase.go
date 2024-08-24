package transform

import "strings"

type LowercaseTransformer struct{}

func NewLowercaseTransformer() *LowercaseTransformer {
	return &LowercaseTransformer{}
}

func (t *LowercaseTransformer) Transform(words []string) []string {
	for i, word := range words {
		words[i] = strings.ToLower(word)
	}
	return words
}
