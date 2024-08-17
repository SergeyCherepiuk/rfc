package utils

import "strings"

func SplitWithWhiteList(text string, whitelist []rune) []string {
	lookup := make(map[rune]struct{}, len(whitelist))
	for _, allowed := range whitelist {
		lookup[allowed] = struct{}{}
	}

	isAllowed := func(r rune) bool {
		_, found := lookup[r]
		return !found
	}

	return strings.FieldsFunc(text, isAllowed)
}
