package utils

func FilterRegularWords(words []string) []string {
	predicates := []func(string) bool{
		predicateZeroOrOneLetterWord,
		predicateUppercaseWord,
		predicateWordWithNumbers,
	}

	wordsKept := 0
	for _, word := range words {
		if !matchesAny(word, predicates) {
			words[wordsKept] = word
			wordsKept++
		}
	}
	return words[:wordsKept]
}

func matchesAny(word string, predicates []func(string) bool) bool {
	for _, predicate := range predicates {
		if predicate(word) {
			return true
		}
	}
	return false
}

func predicateZeroOrOneLetterWord(word string) bool {
	return len(word) <= 1
}

func predicateUppercaseWord(word string) bool {
	for _, letter := range word {
		if letter < 'A' || letter > 'Z' {
			return false
		}
	}
	return true
}

func predicateWordWithNumbers(word string) bool {
	for _, letter := range word {
		if letter >= '0' && letter <= '9' {
			return true
		}
	}
	return false
}
