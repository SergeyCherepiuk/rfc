package transform

import "strings"

type TrimTransformer struct {
	cutset string
}

func NewTrimTransformer(cutset string) *TrimTransformer {
	return &TrimTransformer{cutset}
}

func (t *TrimTransformer) Transform(words []string) []string {
	for i, word := range words {
		words[i] = strings.Trim(word, "-")
	}
	return words
}
