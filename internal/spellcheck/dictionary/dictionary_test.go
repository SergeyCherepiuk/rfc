package dictionary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDictionary(t *testing.T) {
	type ConstructorTestCase struct {
		name     string
		words    []string
		expected map[byte][]string
	}

	constructorTestCases := []ConstructorTestCase{
		{
			name:     "NoWords",
			words:    []string{},
			expected: make(map[byte][]string, 26),
		},
		{
			name:  "SortedWordsSingleWordPerLetter",
			words: []string{"apple", "banana", "cinnamon", "zurek"},
			expected: map[byte][]string{
				'a': []string{"apple"},
				'b': []string{"banana"},
				'c': []string{"cinnamon"},
				'z': []string{"zurek"},
			},
		},
		{
			name:  "SortedWordsMultipleWordsPerLetter",
			words: []string{"apple", "apricots", "avocado", "blueberry", "bread", "butter"},
			expected: map[byte][]string{
				'a': []string{"apple", "apricots", "avocado"},
				'b': []string{"blueberry", "bread", "butter"},
			},
		},
		{
			name:  "UnsortedWordsSingleWordPerLetter",
			words: []string{"cinnamon", "apple", "zurek", "banana"},
			expected: map[byte][]string{
				'a': []string{"apple"},
				'b': []string{"banana"},
				'c': []string{"cinnamon"},
				'z': []string{"zurek"},
			},
		},
		{
			name:  "UnsortedWordsMultipleWordsPerLetter",
			words: []string{"avocado", "apricots", "blueberry", "bread", "apple", "butter"},
			expected: map[byte][]string{
				'a': []string{"apple", "apricots", "avocado"},
				'b': []string{"blueberry", "bread", "butter"},
			},
		},
	}

	for _, tc := range constructorTestCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := dictionary(tc.words).Words
			assert.Equal(t, tc.expected, actual)
		})
	}

	type AllWordsTestCase struct {
		name     string
		words    []string
		expected []string
	}

	allWordsTestCases := []AllWordsTestCase{
		{
			name:     "EmptyDictionary",
			words:    []string{},
			expected: []string{},
		},
		{
			name:     "NonEmptyDictionary",
			words:    []string{"apple", "banana", "cinnamon"},
			expected: []string{"apple", "banana", "cinnamon"},
		},
	}

	for _, tc := range allWordsTestCases {
		t.Run(tc.name, func(t *testing.T) {
			dictionary := dictionary(tc.words)
			actual := dictionary.AllWords()
			assert.Equal(t, tc.expected, actual)
		})
	}

	type ContainsTestCase struct {
		name     string
		words    []string
		target   string
		expected bool
	}

	containsTestCases := []ContainsTestCase{
		{
			name:     "EmptyDictionary",
			words:    []string{},
			target:   "beans",
			expected: false,
		},
		{
			name:     "DoesNotContainTargetWord",
			words:    []string{"apple", "banana", "cinnamon", "zurek"},
			target:   "beans",
			expected: false,
		},
		{
			name:     "DoesNotContainTargetWord",
			words:    []string{"apple", "banana", "cinnamon", "zurek"},
			target:   "banana",
			expected: true,
		},
	}

	for _, tc := range containsTestCases {
		t.Run(tc.name, func(t *testing.T) {
			dictionary := dictionary(tc.words)
			actual := dictionary.Contains(tc.target)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
