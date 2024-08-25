package dictionary

import (
	"fmt"
	"os"
	"slices"

	"github.com/SergeyCherepiuk/rfc/internal/search"
	"github.com/SergeyCherepiuk/rfc/internal/sort"
	"github.com/SergeyCherepiuk/rfc/internal/transform"
)

const dictionaryFile = "/usr/share/dict/words"

type Dictionary struct {
	Words map[byte][]string
}

func NewDictionary() (Dictionary, error) {
	content, err := os.ReadFile(dictionaryFile)
	if err != nil {
		err = fmt.Errorf("failed to load dictionary file from %s: %w", dictionaryFile, err)
		return Dictionary{}, err
	}

	words := transform.NewPipeline(content).
		AddTransformations(transform.DefaultTransformers...).
		Run()

	return dictionary(words), nil
}

func (d Dictionary) AllWords() []string {
	count := 0
	for _, words := range d.Words {
		count += len(words)
	}

	allWords := make([]string, 0, count)
	for _, words := range d.Words {
		allWords = append(allWords, words...)
	}

	return allWords
}

func (d Dictionary) Contains(word string) bool {
	if len(word) == 0 {
		return false
	}

	letter := word[0]
	words, ok := d.Words[letter]
	if !ok {
		return false
	}

	_, found := search.Binary(words, word)
	return found
}

func dictionary(words []string) Dictionary {
	maybeSortAscending(words)
	wordsByLetter := splitByFirstLetter(words)
	return Dictionary{wordsByLetter}
}

func maybeSortAscending(words []string) {
	direction, sorted := sort.IsSorted(words)
	if !sorted {
		slices.Sort(words)
	} else if sorted && direction == sort.Descending {
		slices.Reverse(words)
	}
}

func splitByFirstLetter(words []string) map[byte][]string {
	wordsByLetter := make(map[byte][]string, 26)

	prev := 0
	for i := 0; i < len(words)-1; i++ {
		// NOTE: sc: Empty strings are filtered out by pipeline in constructor.
		// So accessing first letter by index is safe.
		w1Letter, w2Letter := words[i][0], words[i+1][0]
		if w1Letter != w2Letter {
			wordsByLetter[w1Letter] = words[prev : i+1]
			prev = i + 1
		}
	}

	if prev < len(words) {
		letter := words[prev][0]
		wordsByLetter[letter] = words[prev:]
	}

	return wordsByLetter
}
