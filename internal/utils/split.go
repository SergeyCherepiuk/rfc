package utils

import "strings"

func SplitByAny(text string, separators []rune) []string {
	lookup := make(map[rune]struct{}, len(separators))
	for _, separator := range separators {
		lookup[separator] = struct{}{}
	}

	isSeparator := func(r rune) bool {
		_, found := lookup[r]
		return found
	}

	return strings.FieldsFunc(text, isSeparator)
}
