package dictionary

import (
	"fmt"
	"iter"

	"github.com/SergeyCherepiuk/rfc/internal/spellcheck"
	"github.com/SergeyCherepiuk/rfc/internal/spellcheck/diff"
)

const maxEditDistance = 1

type DictionaryChecker struct {
	dictionary Dictionary
}

func NewDictionaryChecker() (DictionaryChecker, error) {
	dictionary, err := NewDictionary()
	if err != nil {
		err = fmt.Errorf("failed to initialize dictionary: %w", err)
		return DictionaryChecker{}, err
	}

	return DictionaryChecker{dictionary}, nil
}

func (c DictionaryChecker) Check(word string) (spellcheck.CheckResult, error) {
	result := spellcheck.CheckResult{
		Word:        word,
		Correct:     c.dictionary.Contains(word),
		Suggestions: make(map[string]int, 0),
	}

	if result.Correct {
		return result, nil
	}

	for _, w := range c.dictionary.AllWords() {
		lengthDifference := abs(len(word) - len(w))
		if lengthDifference > maxEditDistance {
			continue
		}

		distance := diff.EditDistance(word, w)
		if distance <= maxEditDistance {
			result.Suggestions[w] = distance
		}
	}

	return result, nil
}

func (c DictionaryChecker) IncorrectWords(words []string) iter.Seq2[spellcheck.CheckResult, error] {
	return func(yield func(spellcheck.CheckResult, error) bool) {
		for _, word := range words {
			result, err := c.Check(word)

			if result.Correct {
				continue
			}

			if len(result.Suggestions) == 0 {
				continue
			}

			if !yield(result, err) {
				break
			}
		}
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
