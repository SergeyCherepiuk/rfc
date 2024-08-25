package spellcheck

import (
	"fmt"

	"github.com/SergeyCherepiuk/rfc/internal/spellcheck/dictionary"
	"github.com/SergeyCherepiuk/rfc/internal/spellcheck/diff"
)

const maxEditDistance = 1

type DictionaryChecker struct {
	dictionary dictionary.Dictionary
}

func NewDictionaryChecker() (DictionaryChecker, error) {
	dictionary, err := dictionary.NewDictionary()
	if err != nil {
		err = fmt.Errorf("failed to initialize dictionary: %w", err)
		return DictionaryChecker{}, err
	}

	return DictionaryChecker{dictionary}, nil
}

func (c DictionaryChecker) Check(word string) (bool, map[string]int, error) {
	if c.dictionary.Contains(word) {
		return true, make(map[string]int), nil
	}

	suggestions := make(map[string]int)

	for _, w := range c.dictionary.AllWords() {
		lengthDifference := abs(len(word) - len(w))
		if lengthDifference > maxEditDistance {
			continue
		}

		distance := diff.EditDistance(word, w)
		if distance <= maxEditDistance {
			suggestions[w] = distance
		}
	}

	return false, suggestions, nil
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
